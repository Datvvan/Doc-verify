pragma solidity ^0.8.17;

contract StorageContract {
    struct ContractInfo {
        string name;
        string company;
        string identity_card;
        string hash_docs;
    }

    ContractInfo[] contracts;

    function addContract(
        string memory name,
        string memory company,
        string memory identity_card,
        string memory hash_docs
    ) public {
        ContractInfo memory c = ContractInfo(
            name,
            company,
            identity_card,
            hash_docs
        );
        contracts.push(c);
    }

    function getContract(string memory identity_card, string memory company)
        public
        view
        returns (string memory)
    {
        uint256 i;
        for (i = 0; i < contracts.length; i++) {
            ContractInfo memory c = contracts[i];
            if (
                (keccak256(abi.encodePacked(c.identity_card)) ==
                    keccak256(abi.encodePacked(identity_card))) &&
                (keccak256(abi.encodePacked(c.company)) ==
                    keccak256(abi.encodePacked(company)))
            ) {
                return (c.hash_docs);
            }
        }
        return ("Not found");
    }
}
