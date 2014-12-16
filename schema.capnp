@0xd91714aa5f7ee051;
using Go = import "go.capnp";
$Go.package("hare");
$Go.import("testpkg");


struct DocumentCapn { 
   body  @0:   Text; 
} 

struct FactCapn { 
   kView    @0:   UInt64; 
   kDomain  @1:   UInt64; 
   kTarget  @2:   UInt64; 
} 

struct SequencerCapn { 
   mark  @0:   UInt64; 
} 

struct ViewCapn { 
   src  @0:   Text; 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

