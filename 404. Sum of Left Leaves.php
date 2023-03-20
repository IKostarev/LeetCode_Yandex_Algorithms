<?php

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return Integer
     */
    function sumOfLeftLeaves($root) {
        $sum = 0;
        if ($root->left != null) {
            $left = $root->left;
            if (!$left->left && !$left->right) {
                $sum += $left->val;
            }
            $sum += $this->sumOfLeftLeaves($left);
        }
        if ($root->right) {
            $right = $root->right;
            $sum += $this->sumOfLeftLeaves($right);
        }
        return ($sum ? $sum : 0);
    }
}