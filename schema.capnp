@0xf4e89d71282c618d;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct KeyCapn { 
   id  @0:   UInt64; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

