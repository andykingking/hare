@0xbd91bec0532c686c;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct DocumentCapn { 
   body  @0:   Text; 
} 

struct KeyCapn { 
   kInt  @0:   UInt64; 
} 

struct SequencerCapn { 
   index  @0:   KeyCapn; 
} 

struct ViewCapn { 
   src  @0:   Text; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

