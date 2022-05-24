// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.14;

contract UsersStorage {
    struct User {
        uint64 id;
        string name;
    }
    // Users array
    User[] private Users;
    // Mapping user ID to Users[index]
    struct UserIndex {
        uint index;
        bool isExist;
    }
    mapping (uint64 => UserIndex) private UserIdToIndex;

    // Initializing the state variable
    uint randNonce = 0;
    function randomUint64() private returns (uint64) {
        {
        // increase nonce
        randNonce++;
        uint rand = uint(keccak256(abi.encodePacked(block.timestamp, msg.sender, randNonce)));
        return uint64(rand);
        }
    }

    function UsersCreate(string memory name) public {
        uint limiter = 0;
        uint64 id = randomUint64();
        while (UserIdToIndex[id].isExist) {
            id = randomUint64();
            require(limiter < 10);
            limiter++;
        }
        Users.push(User(id, name));
        UserIdToIndex[id] = UserIndex(Users.length - 1, true);
    }

    function UsersGetOne(uint64 id) public view returns (User memory) {
        require(UserIdToIndex[id].isExist);
        return Users[UserIdToIndex[id].index];
    }

    function UsersGetAll() public view returns (User[] memory) {
        return Users;
    }

    function UsersDelete(uint64 id) public {
        require(UserIdToIndex[id].isExist);
        uint index = UserIdToIndex[id].index;
        if (Users.length > 1) {
            User memory last = Users[Users.length - 1];
            Users[index] = last;
            UserIdToIndex[last.id].index = index;
        }
        Users.pop();
        delete UserIdToIndex[id];
    }
}
