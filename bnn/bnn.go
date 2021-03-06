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
    "math"
)

//essentially a link connecting neurons 
type Synapse struct {
    //weight associated with the synapse
    Weight      float32
    //no of the input/output neuron
    In, Out     int
}

//FIXME calculate deltas locally per neuron for BP
//TODO Inputs and Outputs useful for a sparse net
type Neuron struct {
    //activation function
    Activation  string
    //no of inputs and outputs per neuron
    Inps, Outs  []int
    //for calculating deltas
    DeltaTemp   float32
    //neuron's output
    OutVal      float32
}

//TODO extend to support any activation type
func ActivationFunction(x float64) float64{

    return math.Max(0,x) 
}


//inference takes an input image and uses the weights from training  
//FIXME add bias
//FIXME pass array of layers 
func Inference(weights [][]Synapse, input [][]float32, network [][]Neuron) float32{

   var output float32

   //calculate out values for the first layer (i = 0)
   for _, layer := range network {
     for j, neuron := range layer {
        neuron.OutVal += weights[0][j].Weight * input[0][j]
     }
   }

   //use the weights to calculate the output of neurons in hidden layers
   for i, layer := range network {
     for j, neuron := range layer {
        weights[i][j].Weight += weights[i][j].Weight * neuron.OutVal * input[i][j]
     }
   }

   //use the weights to calculate the output of neurons in final layer (i = last)
   i := len(network)
   for _, layer := range network {
     for j, neuron := range layer {
        output += weights[i][j].Weight * neuron.OutVal * input[i][j]
     }
   }
   return output
}


//trains the network of layers based on the input batches
//compares the output based on the test in the dataset 
//TODO add bias and weight distributions as input 
//TODO pass a pointer to the network
func TrainNetwork(image []byte, test []byte, network [][]Neuron) ([][]Synapse, float32){

   var accuracy float32
   var weights [][]Synapse

   //TODO initialise weights using a random function

   //calculate deltas per neuron
   for i := len(network); i >= 0; i-- {
     for j, _ := range network {

       var acc float32 = 0
       for k, _ := range network {
        acc += weights[i+1][k].Weight * network[i][j].DeltaTemp
       }
       network[i][j].DeltaTemp =  acc
     }
   } 

   //calculate new weights and update
   for i, layer := range network {
     for j, neuron := range layer {
        weights[i][j].Weight += weights[i][j].Weight * neuron.OutVal
     }
   }

   return weights, accuracy
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
func NetworkLayer(size int, act string) []Neuron{

  layer := make([]Neuron, size)

  //init the array
  for i, _:= range layer {

    layer[i].Activation = act
  }
  return layer
}
