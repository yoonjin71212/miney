package main
import (
        "encoding/json"
        "encoding/base64"
        "fmt"
        "log"
        crand "crypto/rand"
        "strconv"
        "math"
        "math/big"
        "net/http"
        "io/ioutil"
        "os"
        rand "math/rand"
        "os/exec"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
        "go.mongodb.org/mongo-driver/bson"
        "github.com/gorilla/mux"
        "context"
        "bytes"
)
var ePlace int64
var SERVER_IP = os.Args[1] 
var PORT_LIST = make([]int64,0,100000)
var flag   bool
var authFlag bool = false
var flags MCFlags
var port   string
var portprev string = "60001"
var cursor interface{}
var route *mux.Router
var route_MC *mux.Router
var current []byte
var current_Config []byte 
var buf bytes.Buffer
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
var col *mongo.Collection
var ipCol , UserCol *mongo.Collection
var portInt int64 = 25563
var portIntonePlace int64 = 25563
var ctx context.Context
var tag string
var password string = "a"
var ADMIN    string = "a"
var ADDR string = "http://daegu.yjlee-dev.pe.kr"

type UserInfo struct {
    Username  string `json:"username"`
    Password  string `json:"password"`
}
func TouchFile (name string) {
        file , _ := os.OpenFile(name , os.O_RDONLY|os.O_CREATE , 0644)
        file.Close()
}
type ContainerInfo struct {
	Servername string `json:"servername"`
  Username   string `json:"username"`
  Password   string `json:"password"`
  TAG        string `json:"tag"`
  Serverip string `json:"serverip"`
	Serverport string `json:"serverport"`
}
var INFO ContainerInfo
type MCFlags struct {
//Name
        Server_Name                string `json:"server-name"`

/* This Section Defines Mode of this game.
         Mode Section 

Note : NetMode Section is Different Section */ 
                                /*Defines The Game Mode .
                                  Allowed Values: 
                                  "survival" , 
                                "creative"  ,
                                  "adventure"
                                 */
        GameMode            string   `json:"gamemode"`
                                /*
                                         Prevent the Server from Sending to the Client GameMode 
                                   Other than the GameMode Defined as Mode string
                                 */
        Force_GameMode          bool `json:"force-gamemode"`
                                /* Defines The Difficulty of this Game

                                         Allow Values :
                                         "peaceful" ,
                                         "easy" ,
                                         "normal" , 
                                         "hard"
                                 */
        Difficulty          string   `json:"difficulty"`
                                /*
                                         Selects Cheat Options . It this is enabled , players can
                                         use Command-line Cheats
                                */
        Allow_Cheat             bool `json:"allow-cheat"`

/* This Section Defines NetMode of This Server.
         Netmode Section 
 */
                                /* Defines the Max Players.
                                         Negative Values are not allowed.
                                 */
        Max_Players         int64 `json:"max-players"`
                                /*
                                        If enabled then all connected players must be authenticated
                                        to Xbox Live ;
                                  ( Recommended by MineCraft Developers )
                                 */
        Online_Mode          bool `json:"online-mode"`
                                /*
                                         If Enabled , then Server accepts only the Whitelisted
                                         Players
                                         Whitelist File Location : ./whitelist.json
                                 */
        White_List           bool `json:"white-list"`

/* Port Of This MineCraft Server is Here
Difference of Two Ports :  
Read Here -- Link :
"https://community.fs.com/blog/ipv4-vs-ipv6-whats-the-difference.html"
 */
        Server_Port         int64 `json:"server-port"`
        Server_PortV6       int64 `json:"server-portv6"`

        View_Distance       int64 `json:"view-distance"`

        Tick_Distance       int64 `json:"tick-distance"`
        Player_Idle_TimeOut int64 `json:"player-idle-timeout"`
                                /* Maximum Number of Thread
                                         The Server can use
                                         1 - System Max Value
                                 */
        MAX_Threads         int64 `json:"max-threads"`

/* Advanced Level Options */
        Level_Name         string `json:"level-name"`
        Level_Seed         string `json:"level-seed"`
/*                       */
                                /*
                                         Player Permissions
                                         "visitor"
                                         "member"
                                         "operator" // Admin
                                 */
        Default_Player_Permission_Level              string  `json:"default-player-permission-level"`

/* Advanced Options */
        TexturePack_Required                 bool `json:"texturepack-required"`
        Content_Log_File_Enabled             bool `json:"content-log-file-enabled"`

        Compression_Threshold                    int64  `json:"compression-threshold"`

		Server_Authoritative_Movement                string  `json:"server-authoritative-movement"`
        Player_Movement_Score_Threshold    float64 `json:"player-movement-score-threshold"`
        Player_Movement_Distance_Threshold float64 `json:"player-movement-distance-threshold"`
        Player_Movement_Duration_Threshold_In_ms int64 `json:"player-movement-duration-threshold-in-ms"`
        Correct_Player_Movement              bool `json:"correct-player-movement"`

        Server_Authoritative_Block_Breaking  bool `json:"server-authoritative-block-breaking"`
/*------- END -------*/
}

func RandStringBytes(n int) string {

        seed , _ := crand.Int (crand.Reader , big.NewInt (math.MaxInt64))
        rand.Seed (seed.Int64())

        b := make([]byte, n)
        for i := range b {
                b[i] = letterBytes[rand.Intn(len(letterBytes))]
        }
        return string(b)
}

func check(u string, pw string) bool {
    if (u==ADMIN) && !botCheck(u,pw) {
        return true
    }
    return false
}
func decrypt(pw string) string {
    str, _ := base64.StdEncoding.DecodeString(pw)
    return string(str)
}

func botCheck(u string,pw string) bool {
    cur, _ := UserCol.Find(context.Background(), bson.D{{}})
	      for cur.Next(context.TODO()) {
            current , _ = bson.MarshalExtJSON( cur.Current , false , false )
            var i UserInfo
            json.Unmarshal(current,&i)
            password = decrypt(i.Password)
            if (password ==pw) && (i.Username==u) {
                return false
            }
        
    }
    return true
}



func get_TAG (mydir string) string {

        var err error
        var file *os.File
        file , err = os.OpenFile ( mydir+"/container/latest_access" , os.O_RDWR , os.FileMode (0644))
        if err != nil {
                log.Println ( tag )
        }
	tagRet := "minecraft-"+RandStringBytes(20)
        file.Write([]byte(tagRet))
        file.Close ()
	return tagRet

}

func Generate_ConfigFile ( bytesGenerated []byte ) []byte {




	var err error
  if(ePlace>0) {
      ePlace--
  		flags.Server_Port = PORT_LIST[0]
      portIntonePlace+=3
  } else {
  		flags.Server_Port = portInt
  }
		flags.Server_PortV6 = portInt+1
    INFO.Serverport= strconv.Itoa(int(portInt))
    port= strconv.Itoa(int(portInt))
    bytesGenerated , err = json.Marshal (flags)
    if err != nil {
    		log.Println (err)
   	}
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("{") , []byte  ("") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("}") , []byte("")   , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte(":") , []byte("=")  , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte(",") , []byte("\n") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("_") , []byte("-") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("\"") , []byte("") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("'") , []byte("") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte(")") , []byte("") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("(") , []byte("") , -1 )
        bytesGenerated = bytes.Replace ( bytesGenerated , []byte("NumberLong") , []byte("") , -1 )
				portInt+=3
        PORT_LIST = append(PORT_LIST,portInt)

        return bytesGenerated

}

func CreateConfig (wr http.ResponseWriter , req *http.Request) {


	user, pass, _  := req.BasicAuth()
  INFO.Username = user
  INFO.Password = pass
  
    if botCheck(user,pass) == true {
        wr.Write( []byte("Unauthorized") )
        return
    }
  if(flag) {
      wr.Write([]byte("wait"))
      return
  }
  flag=true
  wr.Header().Set("Content-Type", "application/json; charset=utf-8")
  bytesGenerated , err := ioutil.ReadAll (req.Body)
  bytesGenerated = bytes.Replace( bytesGenerated , []byte("_") , []byte("-") , -1 )
	bytesGenerated = bytes.TrimLeft( bytesGenerated , " \t\n" )
	bytesGenerated = bytes.TrimRight(bytesGenerated, " \t\n" )
  err = json.Unmarshal (bytesGenerated,&flags)
  if err != nil {
          log.Println (err)
  } else {
          log.Println ( "Request Decoded")
  }
  mydir  := "/usr/local/bin/minecraft"
  tag=get_TAG(mydir)
  if(port==portprev) {
      fmt.Fprintf(wr, "Unauthorized")
      return
  }
  log.Println ("/container_creation.sh " + tag + " " + port)
  portprev = port
  bytesGenerated = Generate_ConfigFile (bytesGenerated)

  TouchFile (mydir+"/properties/"+tag+".properties")
  file , _ := os.OpenFile ( mydir+"/properties/"+tag+".properties" , os.O_RDWR  , os.FileMode(0777))
  log.Println ("Using : properties/"+tag+".properties")

  fmt.Fprintf ( file , string (bytesGenerated) )
  file.Close()
  cmdCreate := exec.Command("/bin/bash","-c","container_creation.sh " + tag + " " + port)
  cmdCreate.Stdout = os.Stdout
  cmdCreate.Stderr = os.Stderr
  err = cmdCreate.Start()
  if err != nil {
          log.Println(err)
  }
  cmdCreate.Wait()
  fmt.Println ("using "+tag)


  if len(flags.Server_Name) != 0 {
          INFO.Servername=flags.Server_Name
          INFO.TAG = tag
  } else {
  	INFO.Servername="MineCraft-Server"
  }
  string_Reply , _ := json.Marshal (INFO)
  mcEXEC :=  exec.Command("/bin/bash" ,"-c","init_server.sh "+tag)
  mcEXEC.Stdout = os.Stdout
  mcEXEC.Stderr = os.Stderr
  mcEXEC.Start()
  bytesGenerated = bytes.Replace ( bytesGenerated , []byte("\n") , []byte  ("") , -1 )
  result , insertErr := col.InsertOne(ctx , flags)

  if insertErr != nil {
          log.Println ( "Insert Error : " ,  insertErr )
  } else {
          fmt.Println ( "Insert Succeed. Result is : " , result )
  }

  ipRes , insertErr := ipCol.InsertOne(ctx , INFO)

  if insertErr != nil {
          log.Println ( "Cannot insert container IP into MongoDB")
  } else {
          log.Println ( "container IP Insert succeed. Result is : " , ipRes)
  }

  string_Reply , err = json.Marshal (INFO)


  if flag == true {
  	flag = false
  }
  wr.Write( string_Reply )
}

func UseConfig (wr http.ResponseWriter , req *http.Request) {
        wr.Header().Set("Content-Type", "application/json; charset=utf-8")
        var cursor interface{}
	var err error
        current_Config , err = ioutil.ReadAll( req.Body )
        if err != nil {
                log.Println (err)
                return
        }
        err = bson.UnmarshalExtJSON( current_Config , false , cursor )
        curr , err :=  ipCol.Find( ctx  , cursor , nil )

        if err != nil {
                log.Println (err)
                return
        }

        current , err = bson.MarshalExtJSON( curr , false , false )

        if err != nil {
                log.Println (err)
                current = nil
        } else {
                wr.Write( current )
        }


}
func DeleteFromListByValue(slice []int64, value int64) []int64 {
    for i, itm := range slice {
        if(itm==value) {
            slice = append(slice[:i],slice[i+1:]...)
        }
    }
    return slice
}
func StopByTag (wr http.ResponseWriter, req *http.Request) {
    forTag, _ := ioutil.ReadAll(req.Body)
    stringForStopTask := string(forTag)
    cmdStop := exec.Command("/bin/bash","-c", "stop.sh "+stringForStopTask)
    cmdStop.Start()
    cmdStop.Wait()
    return
}
func StartByTag(wr http.ResponseWriter, req *http.Request) {
    forTag, _ := ioutil.ReadAll(req.Body)
    stringForStartTask := string(forTag)
    cmdStart := exec.Command("/bin/bash","-c", "start.sh "+stringForStartTask)
    cmdStart.Start()
    cmdStart.Wait()
    return
}
func DeleteByTag ( wr http.ResponseWriter , req *http.Request) {

  forTag, _ := ioutil.ReadAll(req.Body)
  stringForTag := string(forTag)
  cmdDelete := exec.Command("/bin/bash","-c", "delete_container.sh "+stringForTag)
  cur, _ := ipCol.Find(context.Background(), bson.D{{}})
	for cur.Next(context.TODO()) {
    resp , _ := bson.MarshalExtJSON ( cur.Current , false , false )
    var INFO ContainerInfo
    json.Unmarshal(resp,&INFO)
    if(INFO.TAG!=stringForTag) {
        p32, _ := strconv.Atoi(INFO.Serverport)
        p      := int64(p32)
        PORT_LIST = DeleteFromListByValue(PORT_LIST,p)
        ipCol.DeleteOne(context.Background(),cur.Current)
        portIntonePlace = p
        ePlace += 1
        cmdDelete = exec.Command("/bin/bash","-c", "add_port.sh",INFO.Serverport,INFO.Serverip)
        cmdDelete.Stdout = os.Stdout 
        cmdDelete.Stderr = os.Stderr
        cmdDelete.Start()
        cmdDelete.Wait()
    }
	}
  cmdDelete.Stdout = os.Stdout 
  cmdDelete.Stderr = os.Stderr
  cmdDelete.Start()
  cmdDelete.Wait()
}
func GetConfig ( wr http.ResponseWriter , req *http.Request) {
	INFO.Serverip = SERVER_IP
  wr.Header().Set("Content-Type", "application/json; charset=utf-8")
  read,_:=ioutil.ReadAll ( req.Body )
	if f, ok := wr.(http.Flusher); ok { 
		f.Flush() 
	}

  var in UserInfo
  json.Unmarshal(read,&in)
	var resp []byte
  cur, err := ipCol.Find(context.Background(), bson.D{{}})
  jsonList := make ([]string , 0 , 100000)
	for cur.Next(context.TODO()) {
		resp , err = bson.MarshalExtJSON ( cur.Current , false , false )
		if err != nil {
		    log.Println (err)
		}
    var info UserInfo
    json.Unmarshal(resp,&info)
    if(info.Username==in.Username && info.Password==in.Password) {
        jsonList = append ( jsonList , string(resp) )
    }

	}
	if err != nil {
	    log.Println (err)
	}
	resp , err = json.Marshal(jsonList)

	if err != nil {
		log.Println (err)
	}

	fmt.Fprintf(wr,string(resp))

}

func Register ( wr http.ResponseWriter , req *http.Request) {
  user, pass, _ := req.BasicAuth()
  pass =base64.StdEncoding.EncodeToString([]byte(pass))
  var u UserInfo
  u.Username = user
  u.Password = pass
  UserCol.InsertOne(ctx , u)
  fmt.Fprintf(wr,"Registered User")
}

func main() {
        route = mux.NewRouter()
        route.HandleFunc  ("/register", Register).Methods("GET")
        route.HandleFunc ( "/create" , CreateConfig).Methods("POST")
        route.HandleFunc ( "/request" ,GetConfig).Methods("POST")
        route.HandleFunc ( "/delete" , DeleteByTag).Methods("POST")
        route.HandleFunc ( "/stop" , StopByTag).Methods("POST")
        route.HandleFunc ( "/start" , StartByTag).Methods("POST")
        clientOptions := options.Client().ApplyURI ("mongodb://localhost:27017")
        client , _ := mongo.Connect (context.TODO() , clientOptions)
        clientIP , _ := mongo.Connect (context.TODO() , clientOptions)
        ctx, _ = context.WithCancel(context.Background())
        col    = client.Database("MC_Json").Collection("Flag Collections")
        ipCol  = clientIP.Database("MC_IP").Collection("IP Collections")
        UserCol  = clientIP.Database("MC_USER").Collection("User Collections")
        log.Println(http.ListenAndServe(":32000" , route))

}

