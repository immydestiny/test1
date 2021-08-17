func (s *Scanner) Detect(osVer string, pkgs []ftypes.Package) ([]types.DetectedVulnerability, error) {
         osVer = "7.6"
        //:= 用于局部变量声明
        pkgs = []ftypes.Package{
               {
                                              Name:       "vim-minimal",
                                              Version:    "7.4.160",
                                              Release:    "5.el7",
                                               Epoch:      2,
                                              Arch:       "x86_64",
                                              SrcName:    "vim",
                                               SrcVersion: "7.4.160",
                                               SrcRelease: "5.el7",
                                               SrcEpoch:   2,
                                               Layer: ftypes.Layer{
                                                       DiffID: "sha256:932da51564135c98a49a34a193d6cd363d8fa4184d957fde16c9d8527b3f3b02",
                                               },

               },

               {
                                                Name:       "nss",
                                                Version:    "3.36.0",
                                                Release:    "7.1.el7_6",
                                                Epoch:      0,
                                                Arch:       "x86_64",
                                                SrcName:    "nss",
                                                SrcVersion: "3.36.0",
                                                SrcRelease: "7.4.160",
                                                SrcEpoch:   0,
                },
                  {
    Release: "1.el7_6",
    Version: "4.0.9.2",
    Arch: "noarch",
    Name: "dnf",
  },
                         {
                             Version:    "7.4.160",
                             Name:       "vim-minimal",
                             Arch:       "x86_64",
                             Release:    "5.el7",
                             Epoch:      2,
                },

                {
                           Name:         "docker",
                           Version: "1.13.1",
                           Release: "102.1.git7f2769b.el7",
                           Epoch: 2,
                           Arch: "x86_64",
                },

{
Name: "NetworkManager",
Version: "1.10.2",
Release: "16.el7_5",
Epoch: 1,
Arch: "x86_64",
SrcName: "NetworkManager",
SrcVersion: "1.10.2",
SrcRelease: "16.el7_5",
SrcEpoch: 1,
},
{
Name:"audit",Version:"2.8.1",Release:"3.el7_5.1",Arch:"x86_64",SrcName:"audit",SrcVersion:"2.8.1",SrcRelease:"3.el7_5.1",
},
{
Name:"audit-libs",Version:"2.8.1",Release:"3.el7_5.1",Arch:"x86_64",SrcName:"audit",SrcVersion:"2.8.1",SrcRelease:"3.el7_5.1",
},


        }
 
