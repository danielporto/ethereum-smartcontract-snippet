// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;
//https://github.com/rocky/solc-vscode/blob/5d85f259550b570d54b47e0e6c093cf9231ecd32/src/test/suite/resources/sort.sol
/** @title Classic QuickSort. See https://blog.cotten.io/thinking-in-solidity-6670c06390a9 */
contract QuickSort {
    struct Set {
        mapping(bytes32 => uint256) arraySize;
        mapping(bytes32 => uint256) hashCounter;
        bytes32[] keyList;
    }

    Set SortedElements;
    
    event PrintArrayInfo(uint256, bytes32, uint256[]); //array size, hash, elements(list)
    event PrintSetElementQty(uint256,bytes32,uint256,uint256,uint256); //array size, hash, index, mapsize, quantity (#operations)
    event PrintArray(uint256 , uint256[]); //size, elements
    event PrintHash(bytes32);
    

/*
* insert keeps track of arrays sorted. If a value is incorrectly added to the storage
* we can detect after the experiment finishes.
*/
   function insert(bytes32 _hash, uint256 _size) internal {
       if (SortedElements.hashCounter[_hash]==0) {
           SortedElements.arraySize[_hash] = _size; 
           SortedElements.keyList.push(_hash); //allow iterate over the maps later
       }
       SortedElements.hashCounter[_hash]++;
   }
   
/*
* printAllData will show how many times an given array was sorted.
* during one single experiment, only one array should be printed. Any other element
* is potentially a byzantine fault.
*
* If experiments with different array sizes run, we can keep track of
* the different arrays sorted.
*
* useful at the end of each experiment.
*/   

   function printAllData() public {
        for(uint i = 0; i < SortedElements.keyList.length; i++ ){
            bytes32 hash = SortedElements.keyList[i];
            emit PrintSetElementQty(SortedElements.arraySize[hash], hash, i, SortedElements.keyList.length, SortedElements.hashCounter[hash]);
        }    
   }



/*
* receives an interger and generate an array in reverse order to be sorted and simulate CPU load
* as in cpuheavy.sol in blockbench.
*/

  function sort(uint size) public {
      
    uint[] memory data = new uint256[](size);
        for (uint x = 0; x < data.length; x++) {
            data[x] = size - x;
        }

    quickSort(data, int(0), int(data.length - 1));
    bytes32 hash = keccak256(abi.encode(data));
    insert(hash, size);
    emit PrintHash(hash);

  }


/*
* vebose version of sort. 
*/
  function debugSort(uint size) public {
      
    uint[] memory data = new uint256[](size);
        for (uint x = 0; x < data.length; x++) {
            data[x] = size - x;
        }

    // before
    bytes32 hash0 = keccak256(abi.encode(data));
    emit PrintArrayInfo(size, hash0, data);

    quickSort(data, int(0), int(data.length - 1));
    bytes32 hash = keccak256(abi.encode(data));
    insert(hash,size);
    emit PrintArrayInfo(size, hash , data);

  }


  
  /** @dev Classic quicksort sorting algorithm. 
    * @param arr array to be sorted
    * @param left left-most index of array items needing. Array items to the left of left are already sorted.
    * @param right right-most index of array of items needing sorting. Array items to the right of right are already sorted.
    *
    */
  function quickSort(uint[] memory arr, int left, int right) internal pure {
    int i = left;
    int j = right;
    if (i==j) return;
    uint pivot = arr[uint(left + (right - left) / 2)];
    while (i <= j) {
      while (arr[uint(i)] < pivot) i++;
      while (pivot < arr[uint(j)]) j--;
      if (i <= j) {
        (arr[uint(i)], arr[uint(j)]) = (arr[uint(j)], arr[uint(i)]);
        i++;
        j--;
      }
    }

    if (left < j)
      quickSort(arr, left, j);
    if (i < right)
      quickSort(arr, i, right);
  }
}