// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

contract Todo {

    Task[] public tasks;

    address public owner;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isOwner() {
        require(owner == msg.sender);
        _;
    }

    function add(string memory _contnet) public isOwner {
        tasks.push(Task(_contnet, false));
    }

    function get(uint _id) public view isOwner returns (Task memory) {
        return tasks[_id];
    }

    function list() public view isOwner returns (Task[] memory) {
        return tasks;
    }

    function update(uint _id, string memory _content) public isOwner {
        tasks[_id].content = _content;
    }

    function remove(uint _id) public isOwner {
        for (uint i = _id; i < tasks.length - 1; i--) {
            tasks[i] = tasks[i + 1];
        }
        tasks.pop();
    }
}
