// Code generated by Thrift Compiler (0.16.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/bippio/go-impala/services/fb303"
)

var _ = fb303.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  string getName()")
  fmt.Fprintln(os.Stderr, "  string getVersion()")
  fmt.Fprintln(os.Stderr, "  fb_status getStatus()")
  fmt.Fprintln(os.Stderr, "  string getStatusDetails()")
  fmt.Fprintln(os.Stderr, "   getCounters()")
  fmt.Fprintln(os.Stderr, "  i64 getCounter(string key)")
  fmt.Fprintln(os.Stderr, "  void setOption(string key, string value)")
  fmt.Fprintln(os.Stderr, "  string getOption(string key)")
  fmt.Fprintln(os.Stderr, "   getOptions()")
  fmt.Fprintln(os.Stderr, "  string getCpuProfile(i32 profileDurationInSec)")
  fmt.Fprintln(os.Stderr, "  i64 aliveSince()")
  fmt.Fprintln(os.Stderr, "  void reinitialize()")
  fmt.Fprintln(os.Stderr, "  void shutdown()")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
  var m map[string]string = h
  return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
  parts := strings.Split(value, ": ")
  if len(parts) != 2 {
    return fmt.Errorf("header should be of format 'Key: Value'")
  }
  h[parts[0]] = parts[1]
  return nil
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  headers := make(httpHeaders)
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  var cfg *thrift.TConfiguration = nil
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
    if len(headers) > 0 {
      httptrans := trans.(*thrift.THttpClient)
      for key, value := range headers {
        httptrans.SetHeader(key, value)
      }
    }
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransportConf(trans, cfg)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := fb303.NewFacebookServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "getName":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetName requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetName(context.Background()))
    fmt.Print("\n")
    break
  case "getVersion":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetVersion requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetVersion(context.Background()))
    fmt.Print("\n")
    break
  case "getStatus":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetStatus requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetStatus(context.Background()))
    fmt.Print("\n")
    break
  case "getStatusDetails":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetStatusDetails requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetStatusDetails(context.Background()))
    fmt.Print("\n")
    break
  case "getCounters":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetCounters requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetCounters(context.Background()))
    fmt.Print("\n")
    break
  case "getCounter":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCounter requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetCounter(context.Background(), value0))
    fmt.Print("\n")
    break
  case "setOption":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetOption requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SetOption(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getOption":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOption requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GetOption(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getOptions":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetOptions requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetOptions(context.Background()))
    fmt.Print("\n")
    break
  case "getCpuProfile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCpuProfile requires 1 args")
      flag.Usage()
    }
    tmp0, err45 := (strconv.Atoi(flag.Arg(1)))
    if err45 != nil {
      Usage()
      return
    }
    argvalue0 := int32(tmp0)
    value0 := argvalue0
    fmt.Print(client.GetCpuProfile(context.Background(), value0))
    fmt.Print("\n")
    break
  case "aliveSince":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "AliveSince requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.AliveSince(context.Background()))
    fmt.Print("\n")
    break
  case "reinitialize":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Reinitialize requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Reinitialize(context.Background()))
    fmt.Print("\n")
    break
  case "shutdown":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "Shutdown requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.Shutdown(context.Background()))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
