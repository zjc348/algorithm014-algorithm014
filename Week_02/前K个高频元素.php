<?php


class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function topKFrequent($nums, $k) {
        $ret= [];
        $arr = array_count_values($nums);
        $queue = new SplPriorityQueue();
        foreach($arr as $key => $val){
            $queue->insert($key, $val);
        }
        for($i = 0;$i < $k; $i++) $ret[] = $queue->extract();
        return $ret;
    }
}
