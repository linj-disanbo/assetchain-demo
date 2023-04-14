 pragma solidity ^0.8.6;

   import "@openzeppelin/contracts/utils/Counters.sol";
   import "@openzeppelin/contracts/access/Ownable.sol";

   contract MyProof is  Ownable {
       using Counters for Counters.Counter;
       Counters.Counter private _proofIds;
        mapping (string => uint256) private _proofs;
        mapping (string => address) private _proofSenders;

       constructor() {}

       function getProofSender(string memory proofHash)
           public view returns (address)
       {
        require(_proofs[proofHash] != 0, "proofHash not exists");
           return _proofSenders[proofHash];
       }

         function saveProof(string memory proofHash)
           public  returns (uint256)
       {
            require(_proofs[proofHash] == 0, "proofHash already exists");
           _proofIds.increment();

           uint256 newId = _proofIds.current();
           _proofs[proofHash] =  newId;
           _proofSenders[proofHash] = msg.sender;
           return newId;
       }

   }
