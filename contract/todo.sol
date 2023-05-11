// SPDX-License-Identifier: GPL-3.0

// set compatible solidity compilier versions
pragma solidity >=0.8.2 <0.9.0;

contract ToDo {
    // define variables and types
    address public owner;
    Task[] tasks;
    
    struct Task {
        string content;
        bool status;
    }

    constructor() {
        // set contract owner to wallet that mints the smart contract
        owner = msg.sender;
    }

    // method modifier to check if the wallet calling smart contract methods is the original minter of the smart contract
    modifier isOwner() {
        require(owner == msg.sender);
        // call the method being modified after we check that the wallet calling the contract is the owner of the contract
        _;
    }

    // method to add tasks, only callable by the owner of the smart contract, determined by method modifier
    // use the memory keyword in order update, or delete the task during runtime of the smart contract
    function add(string memory _content) public isOwner {
        // add a task to the the task list, status defaults to false
        tasks.push(Task(_content, false));
    }

    // method to retrieve tasks by task id
    // use the memory keyword in order update, or delete the task during runtime of the smart contract
    function get(uint _id) public view returns (Task memory) {
        return tasks[_id];
    }
    
    // method to list all tasks
    // use the memory keyword in order update, or delete the task during runtime of the smart contract
    function list() public view returns (Task[] memory) {
        return tasks;
    }

    // method to update a task by id, only callable by the owner of the smart contract, determined by method modifier
    // use the memory keyword in order update, or delete the task during runtime of the smart contract
    function update(uint _id, string memory _content) public isOwner {
        tasks[_id].content = _content;
    }

    // method to update tasks status to true or false, only callable by the owner of the smart contract, determined by method modifier
    function toggle(uint _id) public isOwner {
        tasks[_id].status = !tasks[_id].status;
    }

    // method to remove a task from the tasks list by id
    function remove(uint _id) public isOwner {
        // check to make sure that the id being passed is in the task list
        require(tasks.length > _id, "Out of bounds");

        // iterate over the tasks list, in order to move the task to be deleted to the end of the task list
        for (uint i = _id; i < tasks.length -1; i++) {
            tasks[i] = tasks[i + 1];
        }

        // remove the last task in the task list
        tasks.pop();
    }

}