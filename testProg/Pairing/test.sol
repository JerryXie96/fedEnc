pragma solidity >=0.4.22;
pragma experimental ABIEncoderV2;

contract test{
    struct Data{
        bytes a;
        bytes b;
    }
    
    Data currentData;
    
    bool isFit=false;
    
    function store(bytes memory a, bytes memory b) public{
        currentData.a=a;
        currentData.b=b;
    }
    
    function compare(bytes memory a, bytes memory b) public {
        bytes memory input=abi.encodePacked(currentData.a,currentData.b,a,b);
        uint[1]memory output;
        uint length=input.length;
        assembly {
            if iszero(call(not(0),0x08,0,add(input,0x20),length,output,0x20)){
                
                revert(0,0)
            }
        }
        if(output[0]!=0){
            isFit=true;
        }
    }
    
    function reIsFit() public view returns(bool){
        return isFit;
    }
    
    function storeTest(Data memory d) public{
        currentData=d;
    }
    
}