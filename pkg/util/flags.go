package util

import (
	"github.com/spf13/cobra"
	core "k8s.io/api/core/v1"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions/resource"
)

func AddGetFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", core.NamespaceDefault, "List the requested object(s) from this namespace.")
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.")
	cmd.Flags().Bool("all-namespaces", false, "If present, list the requested object(s) across all namespaces. Namespace specified with --namespace will be ignored.")
	cmd.Flags().Bool("show-kind", false, "If present, list the resource type for the requested object(s).")
	cmd.Flags().StringP("output", "o", "", "Output format. One of: json|yaml|wide|name.")
	cmd.Flags().BoolP("show-all", "a", false, "When printing, show all resources (default hide terminated pods.)")
	cmd.Flags().Bool("show-labels", false, "When printing, show all labels as the last column (default hide labels column)")
}

func AddCreateFlags(cmd *cobra.Command, options *resource.FilenameOptions) {
	cmd.Flags().StringP("namespace", "n", core.NamespaceDefault, "Create object(s) in this namespace.")
	usage := "create the resource"
	AddFilenameOptionFlags(cmd, options, usage)
}

func AddDeleteFlags(cmd *cobra.Command, options *resource.FilenameOptions) {
	cmd.Flags().StringP("namespace", "n", core.NamespaceDefault, "Delete object(s) from this namespace.")
	cmd.Flags().BoolP("all", "", false, "Delete all resources, including uninitialized ones, in the namespace of the specified resource types.")
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on.")
	cmd.Flags().StringP("output", "o", "", "Output mode. Use \"-o name\" for shorter output (resource/name).")
	usage := "delete the resource"
	AddFilenameOptionFlags(cmd, options, usage)
}

func AddDescribeFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", core.NamespaceDefault, "Describe object(s) from this namespace.")
	cmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.")
	cmd.Flags().Bool("all-namespaces", false, "If present, describe the requested object(s) across all namespaces. Namespace specified with --namespace will be ignored.")
}

func AddEditFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", core.NamespaceDefault, "Edit object(s) in this namespace.")
	cmd.Flags().StringP("output", "o", "yaml", "Output format. One of: yaml|json.")
}

func AddFilenameOptionFlags(cmd *cobra.Command, options *resource.FilenameOptions, usage string) {
	cmd.Flags().StringSliceVarP(&options.Filenames, "filename", "f", options.Filenames, "Filename to use to "+usage)
	cmd.Flags().BoolVarP(&options.Recursive, "recursive", "R", options.Recursive, "Process the directory used in -f, --filename recursively.")
}

func GetNamespace(cmd *cobra.Command) (string, bool) {
	return cmdutil.GetFlagString(cmd, "namespace"), cmd.Flags().Changed("namespace")
}

func AddAuditReportFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("namespace", "n", core.NamespaceDefault, "Export summary report of the requested object from this namespace.")
	cmd.Flags().String("operator-namespace", "kube-system", "Name of namespace where operator is running")
	cmd.Flags().String("index", "", "Export summary report for this only.")
	cmd.Flags().String("output", "", "Directory used to store summary report")
}

func AddCompareFlags(cmd *cobra.Command) {
	cmd.Flags().String("output", "", "Directory used to store summary report")
	cmd.Flags().Bool("show", true, "If true, comparison result will be printed.")
}
