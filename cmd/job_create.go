package cmd

import (
	"fmt"
	"terrakube/client/models"

	"github.com/spf13/cobra"
)

var JobCreateExample string = `Create a new job
    %[1]v job create --organization-id e5ad0642-f9b3-48b3-9bf4-35997febe1fb -w e5ad0642-f9b3-48b3-9bf4-35997febe1fb  -c apply`

var JobCreateOrgId string
var JobCreateWorkspaceId string
var JobCreateTemplateReference string

var createJobCmd = &cobra.Command{
	Use:   "create",
	Short: "create a job",
	Run: func(cmd *cobra.Command, args []string) {
		createJob()
	},
	Example: fmt.Sprintf(JobCreateExample, rootCmd.Use),
}

func init() {
	jobCmd.AddCommand(createJobCmd)
	createJobCmd.Flags().StringVarP(&JobCreateTemplateReference, "template-reference", "t", "", "Job template reference: <UUID> (required)")
	_ = createJobCmd.MarkFlagRequired("template-reference")
	createJobCmd.Flags().StringVarP(&JobCreateOrgId, "organization-id", "", "", "Organization Id (required)")
	_ = createJobCmd.MarkFlagRequired("organization-id")
	createJobCmd.Flags().StringVarP(&JobCreateWorkspaceId, "workspace-id", "w", "", "Workspace Id (required)")
	_ = createJobCmd.MarkFlagRequired("workspace-id")
}

func createJob() {
	client := newClient()

	job := models.Job{
		Attributes: &models.JobAttributes{
			TemplateReference: JobCreateTemplateReference,
		},
		Type: "job",
		Relationships: &models.JobRelationships{
			Workspace: &models.JobRelationshipsWorkspace{
				Data: &models.JobRelationshipsWorkspaceData{
					Type: "workspace",
					ID:   JobCreateWorkspaceId,
				},
			},
		},
	}

	resp, err := client.Job.Create(JobCreateOrgId, job)

	if err != nil {
		fmt.Println(err)
		return
	}

	renderOutput(resp, output)
}
