@0xbec2cc1e5e0ac136;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct KeyCapn { 
   digest  @0:   UInt64; 
   index   @1:   UInt64; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

