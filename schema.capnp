@0x9c46ab101b57828e;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct DocumentCapn { 
   body  @0:   Text; 
} 

struct KeyCapn { 
   uint64  @0:   UInt64; 
} 

struct SequencerCapn { 
   index  @0:   Int64; 
} 

struct ViewCapn { 
   src  @0:   Text; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

