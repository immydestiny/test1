package redhat

import (
        "strings"
        "time"
        "io/ioutil"
        //"fmt"
        "encoding/json"

        version "github.com/knqyf263/go-rpm-version"
        "golang.org/x/xerrors"

        "github.com/aquasecurity/fanal/analyzer/os"
        ftypes "github.com/aquasecurity/fanal/types"
        dbTypes "github.com/aquasecurity/trivy-db/pkg/types"
        "github.com/aquasecurity/trivy-db/pkg/vulnsrc/redhat"
        "github.com/aquasecurity/trivy/pkg/log"
        "github.com/aquasecurity/trivy/pkg/scanner/utils"
        "github.com/aquasecurity/trivy/pkg/types"
)



// Detect scans and returns redhat vulenrabilities
func (s *Scanner) Detect(osVer string, pkgs []ftypes.Package) ([]types.DetectedVulnerability, error) {
         osVer = "7.6"
         var data Class
         pkgs = make([]ftypes.Package,0)
//        filename := "/root/test4_data.json"
        filename := "/root/data.json"
        file, err  := ioutil.ReadFile(filename)
        if err != nil {
            log.Logger.Info(" err to read file")
        }
        //fmt.Printf("%v\n", string(file))
        json.Unmarshal(file,&data)
//        fmt.Printf("%v\n", data)
        for _,p := range data.Pkgs{
            pkgs = append(pkgs,p)
        }
        log.Logger.Info(" get pkgs")
//        fmt.Printf("%v\n", pkgs)

       // pkgs = []ftypes.Package{
       //        {
       //                                       Name:       "vim-minimal",
       //                                       Version:    "7.4.160",
       //                                       Release:    "5.el7",
       //                                        Epoch:      2,
       //                                       Arch:       "x86_64",
       //                                       SrcName:    "vim",
       //                                        SrcVersion: "7.4.160",
       //                                        SrcRelease: "5.el7",
       //                                        SrcEpoch:   2,
       //                                        Layer: ftypes.Layer{
       //                                                DiffID: "sha256:932da51564135c98a49a34a193d6cd363d8fa4184d957fde16c9d8527b3f3b02",
       //                                        },
       //        },
       //
       //  {
       //  Name: "NetworkManager",
       //  Version: "1.10.2",
       //  Release: "16.el7_5",
       //  Epoch: 1,
       //  Arch: "x86_64",
       //  SrcName: "NetworkManager",
       //  SrcVersion: "1.10.2",
       //  SrcRelease: "16.el7_5",
       //  SrcEpoch: 1,
       //  },

       //

       // }

        log.Logger.Info("Detecting RHEL/CentOS vulnerabilities...")
        log.Logger.Info("----------------- %s", osVer)
        if strings.Count(osVer, ".") > 0 {
                osVer = osVer[:strings.Index(osVer, ".")]
         //       osVer = "7"
        }

