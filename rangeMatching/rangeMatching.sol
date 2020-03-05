// The smart contract of main experiment in rangeMatching

pragma solidity = 0.5.4;
pragma experimental ABIEncoderV2;

contract rangeMatching{
   struct CipherText{
       bytes c1;
       bytes c2;
   } 
   
   struct IndexItem{
       bytes[2][32] cipher; // the first item is c1 and the second one is c2
       uint id;
   }
   
   struct QueryItem{
       CipherText[2][32] cipher;
   }
   
   IndexItem[] indexCollection;
   
   uint[] result;
   
   uint flag=10;
   
   function store(IndexItem[] memory indexItem) public {
       for(uint i=0;i<indexItem.length;i++){
           indexCollection.push(indexItem[i]);
       }
   }
   
   function search(QueryItem memory query) public {
       uint[] memory equal=new uint[](indexCollection.length);
       bool[] memory isValid=new bool[](indexCollection.length);
       uint pointer=0;
       uint finishedNum=0;
       bytes memory i1;
       bytes memory i2;
       bytes memory q1;
       bytes memory q2;
       // the first scan: match the block 0
       for(uint i=0;i<indexCollection.length;i++){ // for each item in index
           for(uint j=0;j<2;j++){ // for each iteration
               i1=indexCollection[i].cipher[0][0];
               i2=indexCollection[i].cipher[0][1];
               q1=query.cipher[0][j].c1;
               q2=query.cipher[0][j].c2;
               bytes memory input=abi.encodePacked(i1,q2,q1,i2);
               uint[1] memory output;
               uint length=input.length;
               assembly{
                   if iszero(call(not(0),0x08,0,add(input,0x20),length,output,0x20)){
                       revert(0,0)
                   }
               }
               if(i==1 && j==0){
                   flag=output[0];
               }
               if(output[0]!=0 && j==0){ // matched iteration +0: equal
                   equal[pointer]=i;
                   isValid[pointer]=true;
                   pointer++;
               }
               else if (output[0]!=0 && j!=0){ // matched iteration but not +0: matched 
                   result.push(indexCollection[i].id);
                   break;
               }
           }
       }
       
       // scan the number which is the same in block 0
       for(uint i=1;i<32;i++){ // for each block
           if(finishedNum == pointer){ // all the items in equal have been removed
               break;
           }
           for(uint j=0;j<pointer;j++){
               if(isValid[j]==false){ // the current item has been removed
                   continue;
               }
               for(uint k=0;k<2;k++){
                   i1=indexCollection[equal[j]].cipher[i][0];
                   i2=indexCollection[equal[j]].cipher[i][1];
                   q1=query.cipher[i][k].c1;
                   q2=query.cipher[i][k].c2;
                   bytes memory input=abi.encodePacked(i1,q2,q1,i2);
                   uint[1] memory output;
                   uint length=input.length;
                   assembly{
                       if iszero(call(not(0),0x08,0,add(input,0x20),length,output,0x20)){
                           revert(0,0)
                       }
                   }
                   if(output[0]!=0 && k==0){ // still equal
                       break;
                   }
                   else if(output[0]!=0 && k!=0){ // matched
                       result.push(indexCollection[equal[j]].id);
                       isValid[j]=false;
                       break;
                   }
                   else if(output[0]==0 && k==1){ // reach the end of iterations: not matched
                       isValid[j]=false;
                   }
                   
               }
           }
       }
   }
   
   function getResult() public view returns(uint[] memory){
       return result;
   }
   
   function clearResult() public {
       delete result;
   }
}