/*
Copyright 2023 The Rook Authors. All rights reserved.

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

package command

import (
	"github.com/rook/kubectl-rook-ceph/pkg/crds"
	"github.com/rook/kubectl-rook-ceph/pkg/restore"
	"github.com/spf13/cobra"
)

// RestoreCmd represents the restore commands
var RestoreCmd = &cobra.Command{
	Use:     "restore-deleted",
	Short:   "Restores a CR that was accidentally deleted and is still in terminating state.",
	Args:    cobra.RangeArgs(1, 2),
	Example: "kubectl rook-ceph restore-deleted <CRD> [CR_Name]",
	PreRun: func(cmd *cobra.Command, args []string) {
		verifyOperatorPodIsRunning(cmd.Context(), clientSets)
	},
	Run: func(cmd *cobra.Command, args []string) {
		restore.RestoreCrd(cmd.Context(), clientSets, operatorNamespace, cephClusterNamespace, crds.CephRookIoGroup, crds.CephRookResourcesVersion, args)
	},
}
