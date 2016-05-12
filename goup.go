package main

import (
   /*
   #cgo CFLAGS: -I/usr/include/mysql -DMYSQL_DYNAMIC_PLUGIN
   #include <mysql.h>
   #include <string.h>
   #include <stdlib.h>
   #include <stdio.h>
   #include <plugin.h>
   */
   "C"
    "unicode/utf8"
   "strings"
)


//export goup
func goup(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *C.ulong, null_value *C.char, error *C.char) *C.char {
   value := strings.ToUpper(C.GoString(*args.args))

   result = C.CString(string(value))
   *length = C.ulong(utf8.RuneCountInString(value))

   return result;
}


//export goup_init
func goup_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool{
   if(int(args.arg_count) != 1) {
      C.strcpy(message, C.CString("functions requires 1 argument"))
      return 1
   } else {
      return 0
   }
}


func main(){
}
