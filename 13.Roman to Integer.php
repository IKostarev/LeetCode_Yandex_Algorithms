<?php

class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s) {
        $res = 0;
        $i = 0;
        $pre = 0;

        while (isset($s[$i])) {
            switch ($s[$i]) {
                case 'I':
					$map = 1;
					break;

				case 'V':
					$map = 5;
					break;

				case 'X':
					$map = 10;
					break;

				case 'L':
					$map = 50;
					break;

				case 'C':
					$map = 100;
					break;

				case 'D':
					$map = 500;
					break;

				case 'M':
					$map = 1000;
					break;
            }

            $res += $map;

            if ($map > $pre) {
                $res -= 2 * $pre;
            }

            $pre = $map;
            $i++;
        }

        return $res;
        
    }
}