pragma solidity = 0.5.4;
pragma experimental ABIEncoderV2;

contract ExactMatching {
    struct CipherText{
        bytes x;
        bytes y;
    }
    
    struct Index{
        CipherText CT;
        uint id;
    }
    
    Index[] indexCollection;
    
    uint[] result;
    
    function store(CipherText memory ct, uint id) public {
        Index memory indexItem;
        indexItem.CT=ct;
        indexItem.id=id;
        indexCollection.push(indexItem);
    }
    
    function search(CipherText memory ct) public{
        for(uint i=0;i<indexCollection.length;i++){
            bytes memory input=abi.encodePacked(indexCollection[i].CT.x,ct.y,ct.x,indexCollection[i].CT.y);
            uint[1] memory output;
            uint length=input.length;
            assembly{
                if iszero(call(not(0),0x08,0,add(input,0x20),length,output,0x20)){
                    revert(0,0)
                }
            }
            if(output[0]!=0){
                result.push(indexCollection[i].id);
            }
        }
    }
    function getResult() public view returns(uint[] memory){
        return result;
    }
    
    function clearRes() public {
        delete result;
    }
}