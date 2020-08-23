<?php

class MyCircularDeque {
    private $count = 0;
    private $arr = [];
    /**
     * Initialize your data structure here. Set the size of the deque to be k.
     * @param Integer $k
     */
    function __construct($k) {
        $this->count = $k;
    }
  
    /**
     * Adds an item at the front of Deque. Return true if the operation is successful.
     * @param Integer $value
     * @return Boolean
     */
    function insertFront($value) {
        if($this->isFull()){
            return false;
        }
        array_push($this->arr, $value);
        return true;
    }
  
    /**
     * Adds an item at the rear of Deque. Return true if the operation is successful.
     * @param Integer $value
     * @return Boolean
     */
    function insertLast($value) {
        if($this->isFull()){
            return false;
        }
        array_unshift($this->arr, $value);
        return true;
    }
  
    /**
     * Deletes an item from the front of Deque. Return true if the operation is successful.
     * @return Boolean
     */
    function deleteFront() {
        if(array_pop($this->arr) !== null){
            return true;
        }
        
        return false;
    }
  
    /**
     * Deletes an item from the rear of Deque. Return true if the operation is successful.
     * @return Boolean
     */
    function deleteLast() {
        if(array_shift($this->arr) !== null){
            return true;
        }
        return false;
    }
  
    /**
     * Get the front item from the deque.
     * @return Integer
     */
    function getFront() {
        if($this->isEmpty()){
            return -1;
        }
        return $this->arr[count($this->arr) - 1];
    }
  
    /**
     * Get the last item from the deque.
     * @return Integer
     */
    function getRear() {
        if($this->isEmpty()){
            return -1;
        }
        return $this->arr[0];
    }
    /**
     * Checks whether the circular deque is empty or not.
     * @return Boolean
     */
    function isEmpty() {
        if($this->arr){
            return false;
        }
        return true;
    }
  
    /**
     * Checks whether the circular deque is full or not.
     * @return Boolean
     */
    function isFull() {
        return (count($this->arr) >= $this->count)?true:false;
    }
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * $obj = MyCircularDeque($k);
 * $ret_1 = $obj->insertFront($value);
 * $ret_2 = $obj->insertLast($value);
 * $ret_3 = $obj->deleteFront();
 * $ret_4 = $obj->deleteLast();
 * $ret_5 = $obj->getFront();
 * $ret_6 = $obj->getRear();
 * $ret_7 = $obj->isEmpty();
 * $ret_8 = $obj->isFull();
 */