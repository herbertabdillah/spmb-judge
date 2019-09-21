# spmb-judge documentation (stil on process)
Highly customizeable library for selection in college admission, school admission, etc.
(still on process)
(algorithm stil "brute force")
(code not too clean)

## Overview
### Input
- List of department
- List of participant
- Parameter(ex : participant score, department minimum score, department cluster, department quota, etc)
### Output
- List of accepted participant in department
- List of not accepted participant

## Usage
```go
import "github.com/herbertabdillah/spmb-judge"
var s selection.SpmbJudgeContract = selection.NewSelection()
```
## Contract

```go
SetScoringMethod(f func(participantParams *interface{}, args interface{}) float32)
AddDepartment(id int32, quota int32, params interface{})
AddParticipant(id int32, param interface{}, chosenDepartment []int32)
Execute()
GetResult()

```

## Example
### Departement with cluster
- Ipa (science) cluster with parameter = Ipa score
- Ipa (social) cluster with parameter = Ips score
- Departement with Ipa cluster = 0,1. Ips cluster = 2

```go
const(
	Ipa = 0
	Ips = 1
 )
f := func(participantParams *interface{}, args interface{}) float32 {
  var score float32
  param := (*participantParams).(map[string]interface{})
  switch args.(int) {
  case Ipa:
    score = param["ipa"].(float32)
  case Ips:
    score = param["ips"].(float32)
  }
  return score
}
var s selection.SpmbJudgeContract = selection.NewSelection()
s.SetScoringMethod(f)
s.AddDepartment(0,2, Ipa)
s.AddDepartment(1, 2, Ipa)
s.AddDepartment(2, 3, Ips)
s.AddParticipant(0, map[string]interface{}{"ipa" : float32(90), "ips" : float32(80)}, []int32{0,2})
s.AddParticipant(1, map[string]interface{}{"ipa" : float32(85), "ips" : float32(80)}, []int32{2,0,1})
s.AddParticipant(2, map[string]interface{}{"ipa" : float32(80), "ips" : float32(90)}, []int32{0,1})
s.AddParticipant(3, map[string]interface{}{"ipa" : float32(100), "ips" : float32(90)}, []int32{0,1})
s.Execute()
s.GetResult()
 ```
### Departement without cluster. Just use one parameter (score)
```go
f := func(participantParams *interface{}, args interface{}) float32 {  
  score := (*participantParams).(float32)
  return score
}
var s selection.SpmbJudgeContract = selection.NewSelection()
s.SetScoringMethod(f)
s.AddDepartment(0,2, nil)
s.AddDepartment(1, 2, nil)
s.AddDepartment(2, 3, nil)
s.AddParticipant(0, float32(90), []int32{0,2})
s.AddParticipant(1, float32(85), []int32{2,0,1})
s.AddParticipant(2, float32(80), []int32{0,1})
s.AddParticipant(3, float32(100), []int32{0,1})
s.Execute()
s.GetResult()
```
