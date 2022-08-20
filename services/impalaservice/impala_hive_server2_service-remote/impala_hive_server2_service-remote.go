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
	"github.com/bippio/go-impala/services/status"
	"github.com/bippio/go-impala/services/beeswax"
	"github.com/bippio/go-impala/services/cli_service"
	"github.com/bippio/go-impala/services/impalaservice"
)

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__
var _ = impalaservice.GoUnusedProtection__

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TStatus ResetCatalog()")
  fmt.Fprintln(os.Stderr, "  TOpenSessionResp OpenSession(TOpenSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseSessionResp CloseSession(TCloseSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TGetInfoResp GetInfo(TGetInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TExecuteStatementResp ExecuteStatement(TExecuteStatementReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTypeInfoResp GetTypeInfo(TGetTypeInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCatalogsResp GetCatalogs(TGetCatalogsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetSchemasResp GetSchemas(TGetSchemasReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTablesResp GetTables(TGetTablesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTableTypesResp GetTableTypes(TGetTableTypesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetColumnsResp GetColumns(TGetColumnsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetFunctionsResp GetFunctions(TGetFunctionsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetOperationStatusResp GetOperationStatus(TGetOperationStatusReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelOperationResp CancelOperation(TCancelOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseOperationResp CloseOperation(TCloseOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TGetResultSetMetadataResp GetResultSetMetadata(TGetResultSetMetadataReq req)")
  fmt.Fprintln(os.Stderr, "  TFetchResultsResp FetchResults(TFetchResultsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetDelegationTokenResp GetDelegationToken(TGetDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelDelegationTokenResp CancelDelegationToken(TCancelDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TRenewDelegationTokenResp RenewDelegationToken(TRenewDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TGetLogResp GetLog(TGetLogReq req)")
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
  client := impalaservice.NewImpalaHiveServer2ServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "ResetCatalog":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "ResetCatalog requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.ResetCatalog(context.Background()))
    fmt.Print("\n")
    break
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
      flag.Usage()
    }
    arg81 := flag.Arg(1)
    mbTrans82 := thrift.NewTMemoryBufferLen(len(arg81))
    defer mbTrans82.Close()
    _, err83 := mbTrans82.WriteString(arg81)
    if err83 != nil {
      Usage()
      return
    }
    factory84 := thrift.NewTJSONProtocolFactory()
    jsProt85 := factory84.GetProtocol(mbTrans82)
    argvalue0 := cli_service.NewTOpenSessionReq()
    err86 := argvalue0.Read(context.Background(), jsProt85)
    if err86 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseSession requires 1 args")
      flag.Usage()
    }
    arg87 := flag.Arg(1)
    mbTrans88 := thrift.NewTMemoryBufferLen(len(arg87))
    defer mbTrans88.Close()
    _, err89 := mbTrans88.WriteString(arg87)
    if err89 != nil {
      Usage()
      return
    }
    factory90 := thrift.NewTJSONProtocolFactory()
    jsProt91 := factory90.GetProtocol(mbTrans88)
    argvalue0 := cli_service.NewTCloseSessionReq()
    err92 := argvalue0.Read(context.Background(), jsProt91)
    if err92 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetInfo requires 1 args")
      flag.Usage()
    }
    arg93 := flag.Arg(1)
    mbTrans94 := thrift.NewTMemoryBufferLen(len(arg93))
    defer mbTrans94.Close()
    _, err95 := mbTrans94.WriteString(arg93)
    if err95 != nil {
      Usage()
      return
    }
    factory96 := thrift.NewTJSONProtocolFactory()
    jsProt97 := factory96.GetProtocol(mbTrans94)
    argvalue0 := cli_service.NewTGetInfoReq()
    err98 := argvalue0.Read(context.Background(), jsProt97)
    if err98 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ExecuteStatement":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteStatement requires 1 args")
      flag.Usage()
    }
    arg99 := flag.Arg(1)
    mbTrans100 := thrift.NewTMemoryBufferLen(len(arg99))
    defer mbTrans100.Close()
    _, err101 := mbTrans100.WriteString(arg99)
    if err101 != nil {
      Usage()
      return
    }
    factory102 := thrift.NewTJSONProtocolFactory()
    jsProt103 := factory102.GetProtocol(mbTrans100)
    argvalue0 := cli_service.NewTExecuteStatementReq()
    err104 := argvalue0.Read(context.Background(), jsProt103)
    if err104 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteStatement(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTypeInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTypeInfo requires 1 args")
      flag.Usage()
    }
    arg105 := flag.Arg(1)
    mbTrans106 := thrift.NewTMemoryBufferLen(len(arg105))
    defer mbTrans106.Close()
    _, err107 := mbTrans106.WriteString(arg105)
    if err107 != nil {
      Usage()
      return
    }
    factory108 := thrift.NewTJSONProtocolFactory()
    jsProt109 := factory108.GetProtocol(mbTrans106)
    argvalue0 := cli_service.NewTGetTypeInfoReq()
    err110 := argvalue0.Read(context.Background(), jsProt109)
    if err110 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTypeInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetCatalogs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCatalogs requires 1 args")
      flag.Usage()
    }
    arg111 := flag.Arg(1)
    mbTrans112 := thrift.NewTMemoryBufferLen(len(arg111))
    defer mbTrans112.Close()
    _, err113 := mbTrans112.WriteString(arg111)
    if err113 != nil {
      Usage()
      return
    }
    factory114 := thrift.NewTJSONProtocolFactory()
    jsProt115 := factory114.GetProtocol(mbTrans112)
    argvalue0 := cli_service.NewTGetCatalogsReq()
    err116 := argvalue0.Read(context.Background(), jsProt115)
    if err116 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCatalogs(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetSchemas":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSchemas requires 1 args")
      flag.Usage()
    }
    arg117 := flag.Arg(1)
    mbTrans118 := thrift.NewTMemoryBufferLen(len(arg117))
    defer mbTrans118.Close()
    _, err119 := mbTrans118.WriteString(arg117)
    if err119 != nil {
      Usage()
      return
    }
    factory120 := thrift.NewTJSONProtocolFactory()
    jsProt121 := factory120.GetProtocol(mbTrans118)
    argvalue0 := cli_service.NewTGetSchemasReq()
    err122 := argvalue0.Read(context.Background(), jsProt121)
    if err122 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSchemas(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    arg123 := flag.Arg(1)
    mbTrans124 := thrift.NewTMemoryBufferLen(len(arg123))
    defer mbTrans124.Close()
    _, err125 := mbTrans124.WriteString(arg123)
    if err125 != nil {
      Usage()
      return
    }
    factory126 := thrift.NewTJSONProtocolFactory()
    jsProt127 := factory126.GetProtocol(mbTrans124)
    argvalue0 := cli_service.NewTGetTablesReq()
    err128 := argvalue0.Read(context.Background(), jsProt127)
    if err128 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTables(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTableTypes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTableTypes requires 1 args")
      flag.Usage()
    }
    arg129 := flag.Arg(1)
    mbTrans130 := thrift.NewTMemoryBufferLen(len(arg129))
    defer mbTrans130.Close()
    _, err131 := mbTrans130.WriteString(arg129)
    if err131 != nil {
      Usage()
      return
    }
    factory132 := thrift.NewTJSONProtocolFactory()
    jsProt133 := factory132.GetProtocol(mbTrans130)
    argvalue0 := cli_service.NewTGetTableTypesReq()
    err134 := argvalue0.Read(context.Background(), jsProt133)
    if err134 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTableTypes(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetColumns":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetColumns requires 1 args")
      flag.Usage()
    }
    arg135 := flag.Arg(1)
    mbTrans136 := thrift.NewTMemoryBufferLen(len(arg135))
    defer mbTrans136.Close()
    _, err137 := mbTrans136.WriteString(arg135)
    if err137 != nil {
      Usage()
      return
    }
    factory138 := thrift.NewTJSONProtocolFactory()
    jsProt139 := factory138.GetProtocol(mbTrans136)
    argvalue0 := cli_service.NewTGetColumnsReq()
    err140 := argvalue0.Read(context.Background(), jsProt139)
    if err140 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetColumns(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetFunctions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFunctions requires 1 args")
      flag.Usage()
    }
    arg141 := flag.Arg(1)
    mbTrans142 := thrift.NewTMemoryBufferLen(len(arg141))
    defer mbTrans142.Close()
    _, err143 := mbTrans142.WriteString(arg141)
    if err143 != nil {
      Usage()
      return
    }
    factory144 := thrift.NewTJSONProtocolFactory()
    jsProt145 := factory144.GetProtocol(mbTrans142)
    argvalue0 := cli_service.NewTGetFunctionsReq()
    err146 := argvalue0.Read(context.Background(), jsProt145)
    if err146 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetFunctions(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetOperationStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOperationStatus requires 1 args")
      flag.Usage()
    }
    arg147 := flag.Arg(1)
    mbTrans148 := thrift.NewTMemoryBufferLen(len(arg147))
    defer mbTrans148.Close()
    _, err149 := mbTrans148.WriteString(arg147)
    if err149 != nil {
      Usage()
      return
    }
    factory150 := thrift.NewTJSONProtocolFactory()
    jsProt151 := factory150.GetProtocol(mbTrans148)
    argvalue0 := cli_service.NewTGetOperationStatusReq()
    err152 := argvalue0.Read(context.Background(), jsProt151)
    if err152 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetOperationStatus(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelOperation requires 1 args")
      flag.Usage()
    }
    arg153 := flag.Arg(1)
    mbTrans154 := thrift.NewTMemoryBufferLen(len(arg153))
    defer mbTrans154.Close()
    _, err155 := mbTrans154.WriteString(arg153)
    if err155 != nil {
      Usage()
      return
    }
    factory156 := thrift.NewTJSONProtocolFactory()
    jsProt157 := factory156.GetProtocol(mbTrans154)
    argvalue0 := cli_service.NewTCancelOperationReq()
    err158 := argvalue0.Read(context.Background(), jsProt157)
    if err158 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseOperation requires 1 args")
      flag.Usage()
    }
    arg159 := flag.Arg(1)
    mbTrans160 := thrift.NewTMemoryBufferLen(len(arg159))
    defer mbTrans160.Close()
    _, err161 := mbTrans160.WriteString(arg159)
    if err161 != nil {
      Usage()
      return
    }
    factory162 := thrift.NewTJSONProtocolFactory()
    jsProt163 := factory162.GetProtocol(mbTrans160)
    argvalue0 := cli_service.NewTCloseOperationReq()
    err164 := argvalue0.Read(context.Background(), jsProt163)
    if err164 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetResultSetMetadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultSetMetadata requires 1 args")
      flag.Usage()
    }
    arg165 := flag.Arg(1)
    mbTrans166 := thrift.NewTMemoryBufferLen(len(arg165))
    defer mbTrans166.Close()
    _, err167 := mbTrans166.WriteString(arg165)
    if err167 != nil {
      Usage()
      return
    }
    factory168 := thrift.NewTJSONProtocolFactory()
    jsProt169 := factory168.GetProtocol(mbTrans166)
    argvalue0 := cli_service.NewTGetResultSetMetadataReq()
    err170 := argvalue0.Read(context.Background(), jsProt169)
    if err170 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultSetMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "FetchResults":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "FetchResults requires 1 args")
      flag.Usage()
    }
    arg171 := flag.Arg(1)
    mbTrans172 := thrift.NewTMemoryBufferLen(len(arg171))
    defer mbTrans172.Close()
    _, err173 := mbTrans172.WriteString(arg171)
    if err173 != nil {
      Usage()
      return
    }
    factory174 := thrift.NewTJSONProtocolFactory()
    jsProt175 := factory174.GetProtocol(mbTrans172)
    argvalue0 := cli_service.NewTFetchResultsReq()
    err176 := argvalue0.Read(context.Background(), jsProt175)
    if err176 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.FetchResults(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDelegationToken requires 1 args")
      flag.Usage()
    }
    arg177 := flag.Arg(1)
    mbTrans178 := thrift.NewTMemoryBufferLen(len(arg177))
    defer mbTrans178.Close()
    _, err179 := mbTrans178.WriteString(arg177)
    if err179 != nil {
      Usage()
      return
    }
    factory180 := thrift.NewTJSONProtocolFactory()
    jsProt181 := factory180.GetProtocol(mbTrans178)
    argvalue0 := cli_service.NewTGetDelegationTokenReq()
    err182 := argvalue0.Read(context.Background(), jsProt181)
    if err182 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelDelegationToken requires 1 args")
      flag.Usage()
    }
    arg183 := flag.Arg(1)
    mbTrans184 := thrift.NewTMemoryBufferLen(len(arg183))
    defer mbTrans184.Close()
    _, err185 := mbTrans184.WriteString(arg183)
    if err185 != nil {
      Usage()
      return
    }
    factory186 := thrift.NewTJSONProtocolFactory()
    jsProt187 := factory186.GetProtocol(mbTrans184)
    argvalue0 := cli_service.NewTCancelDelegationTokenReq()
    err188 := argvalue0.Read(context.Background(), jsProt187)
    if err188 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "RenewDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RenewDelegationToken requires 1 args")
      flag.Usage()
    }
    arg189 := flag.Arg(1)
    mbTrans190 := thrift.NewTMemoryBufferLen(len(arg189))
    defer mbTrans190.Close()
    _, err191 := mbTrans190.WriteString(arg189)
    if err191 != nil {
      Usage()
      return
    }
    factory192 := thrift.NewTJSONProtocolFactory()
    jsProt193 := factory192.GetProtocol(mbTrans190)
    argvalue0 := cli_service.NewTRenewDelegationTokenReq()
    err194 := argvalue0.Read(context.Background(), jsProt193)
    if err194 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RenewDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetLog":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetLog requires 1 args")
      flag.Usage()
    }
    arg195 := flag.Arg(1)
    mbTrans196 := thrift.NewTMemoryBufferLen(len(arg195))
    defer mbTrans196.Close()
    _, err197 := mbTrans196.WriteString(arg195)
    if err197 != nil {
      Usage()
      return
    }
    factory198 := thrift.NewTJSONProtocolFactory()
    jsProt199 := factory198.GetProtocol(mbTrans196)
    argvalue0 := cli_service.NewTGetLogReq()
    err200 := argvalue0.Read(context.Background(), jsProt199)
    if err200 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetLog(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
