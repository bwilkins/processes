package processes

import (
  "regexp"
  "strconv"
  "os/exec"
)

type PsEntry struct {
  User string
  Pid  int
  CpuUsagePct float64
  MemUsagePct float64
  VirtualMem  int
  ResidentMem int
  TT          string
  State       string
  Started     string
  RunningTime string
  Command     string
}

type PsList []PsEntry

func newline() *regexp.Regexp {
  nl, _ := regexp.Compile(`(?m)\n`)
  return nl
}

func spaces() *regexp.Regexp {
  sp, _ := regexp.Compile(`\s+`)
  return sp
}

func trimSpace(input string) string {
  surround_sp, _ := regexp.Compile(`^\s*(.*)\s*$`)
  return surround_sp.ReplaceAllString(input, "$1")
}


func Ps() PsList {
  ps_cmd := exec.Command("ps", "aux")
  ps_b, _ := ps_cmd.Output()
  ps_output := string(ps_b)
  lines := newline().Split(trimSpace(ps_output), -1)[1:]

  ps_ret := make([]PsEntry, 0, 10)

  var split []string
  var user, tt, state, started, runtime, command string
  var pid, virt, res int
  var cpu, mem float64

  for _,line := range lines {
    if len(line) == 0 {
      continue
    }
    split      = spaces().Split(line, 11)

    user       = split[0]
    pid,     _ = strconv.Atoi(split[1])
    cpu,     _ = strconv.ParseFloat(split[2], 64)
    mem,     _ = strconv.ParseFloat(split[3], 64)
    virt,    _ = strconv.Atoi(split[4])
    res,     _ = strconv.Atoi(split[5])
    tt         = split[6]
    state      = split[7]
    started    = split[8]
    runtime    = split[9]
    command    = split[10]

    ps_ret = append(ps_ret, PsEntry{user, pid, cpu, mem, virt, res, tt, state, started, runtime, command})
  }

  return ps_ret
}
