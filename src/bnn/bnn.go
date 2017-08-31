/*
Copyright 2017 Reconfigure.io Ltd. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package bnn

	
import (
    "os"
)

type neuron struct {
    //activation function
    act string
    //no of inputs and outputs per neuron
    inps, outs int
}



//trains the network of layers based on the input batches
//compares the output based on the test in the dataset 
//FIXME add bias and weight distributions as input 
func TrainNetwork(image []byte, test []byte, layers []neuron) float32{
 
 return accuracy
}

//reshapes images based on the resize factors should support:
//padding, flipping, rotation, transpose, etc.
//FIXME resize to be implemented as a struct wrt alignment factors
//FIXME implement it as a separate package 
func ReshapeImage(image []byte) []byte{
 return image
}

//reads in images located in 'path' and returns an array 
func ReadImage(path string) []byte{

   //open the image file
   f, err := os.Open(path)
   if err != nil {
        panic(err)
   }

   //get the file status
   fi, err := f.Stat()
   if err != nil {
       panic(err)
   }

   //create an arraye of size 'image'
   arr := make([]byte, fi.Size())
   f.Read(arr)

   f.Close()
   return arr
}

//constructs a layer of neurons with arbitrary 'size' and 'activation' functions
func NetworkLayer(size int, act string) []neuron{

  layer := make([]neuron, size)

  //init the array
  for i, _:= range layer {

    layer[i].act = act
    layer[i].inps = 0
    layer[i].outs = 0
  }

  return layer
}


