<?php

/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution {

    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers($l1, $l2) {
        return $this->addTwoNumbersWithCarry($l1, $l2);
    }

    public function addTwoNumbersWithCarry(?ListNode $l1, ?ListNode $l2, int $carry = 0): ListNode
    {
        $x = $y = 0;
		$n1 = $n2 = null;

        if ($l1) {
            $x = $l1->val;
            $n1 = $l1->next;
        }

        if ($l2) {
            $y = $l2->val;
            $n2 = $l2->next;
        }

        $sum = $x + $y + $carry;
        $carry = $sum > 9 ? 1 : 0;

        $node = new ListNode($sum % 10);

        if ($n1 || $n2 || $carry) {
            $node->next = $this->addTwoNumbersWithCarry($n1, $n2, $carry);
        }

        return $node;
    }
}