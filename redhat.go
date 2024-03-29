package redhat

import (
	"strings"
	"time"

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

var (
	redhatEOLDates = map[string]time.Time{
		"4": time.Date(2017, 5, 31, 23, 59, 59, 0, time.UTC),
		"5": time.Date(2020, 11, 30, 23, 59, 59, 0, time.UTC),
		"6": time.Date(2024, 6, 30, 23, 59, 59, 0, time.UTC),
		// N/A
		"7": time.Date(3000, 1, 1, 23, 59, 59, 0, time.UTC),
		"8": time.Date(3000, 1, 1, 23, 59, 59, 0, time.UTC),
	}
	centosEOLDates = map[string]time.Time{
		"3": time.Date(2010, 10, 31, 23, 59, 59, 0, time.UTC),
		"4": time.Date(2012, 2, 29, 23, 59, 59, 0, time.UTC),
		"5": time.Date(2017, 3, 31, 23, 59, 59, 0, time.UTC),
		"6": time.Date(2020, 11, 30, 23, 59, 59, 0, time.UTC),
		"7": time.Date(2024, 6, 30, 23, 59, 59, 0, time.UTC),
		"8": time.Date(2021, 12, 31, 23, 59, 59, 0, time.UTC),
	}
	excludedVendorsSuffix = []string{
		".remi",
	}
)


// Scanner implements the Redhat scanner
type Scanner struct {
	vs dbTypes.VulnSrc
}

// NewScanner is the factory method for Scanner
func NewScanner() *Scanner {
	return &Scanner{
		vs: redhat.NewVulnSrc(),
	}
}


// Detect scans and returns redhat vulenrabilities
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

	log.Logger.Info("Detecting RHEL/CentOS vulnerabilities...")
        log.Logger.Info("----------------- %s", osVer)
	if strings.Count(osVer, ".") > 0 {
		osVer = osVer[:strings.Index(osVer, ".")]
         //       osVer = "7"
	}
	log.Logger.Debugf("redhat: os version: %s", osVer)
	log.Logger.Debugf("redhat: the number of packages: %d", len(pkgs))

	var vulns []types.DetectedVulnerability
	for _, pkg := range pkgs {
		if !s.isFromSupportedVendor(pkg) {
			log.Logger.Debugf("Skipping %s: unsupported vendor", pkg.Name)
			continue
		}

		// For Red Hat Security Data API containing only source package names
		pkgName := addModularNamespace(pkg.SrcName, pkg.Modularitylabel)
		advisories, err := s.vs.Get(osVer, pkgName)
		if err != nil {
			return nil, xerrors.Errorf("failed to get Red Hat advisories: %w", err)
		}

		installed := utils.FormatVersion(pkg)
		installedVersion := version.NewVersion(installed)

		for _, adv := range advisories {
			if adv.FixedVersion != "" {
				continue
			}
			vuln := types.DetectedVulnerability{
				VulnerabilityID:  adv.VulnerabilityID,
				PkgName:          pkg.Name,
				InstalledVersion: installed,
				Layer:            pkg.Layer,
			}
			vulns = append(vulns, vuln)
		}

		// For Red Hat OVAL v2 containing only binary package names
		pkgName = addModularNamespace(pkg.Name, pkg.Modularitylabel)
		advisories, err = s.vs.Get(osVer, pkgName)
		if err != nil {
			return nil, xerrors.Errorf("failed to get Red Hat advisories: %w", err)
		}

		for _, adv := range advisories {
			fixedVersion := version.NewVersion(adv.FixedVersion)
			if installedVersion.LessThan(fixedVersion) {
				vuln := types.DetectedVulnerability{
					VulnerabilityID:  adv.VulnerabilityID,
					PkgName:          pkg.Name,
					InstalledVersion: installed,
					FixedVersion:     fixedVersion.String(),
					Layer:            pkg.Layer,
				}
				vulns = append(vulns, vuln)
			}
		}
	}
	return vulns, nil
}

// IsSupportedVersion checks is OSFamily can be scanned with Redhat scanner
func (s *Scanner) IsSupportedVersion(osFamily, osVer string) bool {
	now := time.Now()
	return s.isSupportedVersion(now, osFamily, osVer)
}

func (s *Scanner) isSupportedVersion(now time.Time, osFamily, osVer string) bool {
	if strings.Count(osVer, ".") > 0 {
		osVer = osVer[:strings.Index(osVer, ".")]
	}

	var eolDate time.Time
	var ok bool
	if osFamily == os.RedHat {
		eolDate, ok = redhatEOLDates[osVer]
	} else if osFamily == os.CentOS {
		eolDate, ok = centosEOLDates[osVer]
	}
	if !ok {
		log.Logger.Warnf("This OS version is not on the EOL list: %s %s", osFamily, osVer)
		return false
	}
	return now.Before(eolDate)
}

func (s *Scanner) isFromSupportedVendor(pkg ftypes.Package) bool {
	for _, s := range excludedVendorsSuffix {
		if strings.HasSuffix(pkg.Release, s) {
			return false
		}
	}
	return true
}

func addModularNamespace(name, label string) string {
	// e.g. npm, nodejs:12:8030020201124152102:229f0a1c => nodejs:12::npm
	var count int
	for i, r := range label {
		if r == ':' {
			count++
		}
		if count == 2 {
			return label[:i] + "::" + name
		}
	}
	return name
}
