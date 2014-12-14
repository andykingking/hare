@0xa1291fe5eb2b2264;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct DocumentCapn { 
   body  @0:   Text; 
} 

struct SequencerCapn { 
   index  @0:   UInt64; 
} 

struct ViewCapn { 
   src  @0:   Text; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

