/**------------------------------------------------------------**
 * @filename cmd/commands.go
 * @author   jiny - caojingyin@jinycoo.com
 * @version  1.0.0
 * @date     2017/07/01 13:17
 * @desc     cmd-commands - summary
 **------------------------------------------------------------**/

package cmd

import "github.com/spf13/cobra"

func AddCommands(cmd *cobra.Command) {
	addVersion(cmd)
	addCreate(cmd)
}
