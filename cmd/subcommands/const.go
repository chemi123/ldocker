package subcommands

// For images subcommand
const (
	shortImagesDesc = "Manages Docker images, including listing and removing images."

	longImagesDesc = `The docker images command provides a list of images available on your local system. This command is essential for managing and understanding the Docker images that are stored locally, which can include images you have built yourself, as well as images pulled from remote Docker registries.

With this command, you can see a comprehensive list of images, along with details such as the repository, tag, image ID, creation time, and size. This visibility is crucial for Docker image management tasks such as cleaning up unused images or identifying the versions of images that are currently available for use.
	
The docker images command supports various options to filter and format the output, making it easier to find specific images or to retrieve the information you need. For example, you can list all images, including intermediate image layers, by using the -a or --all flag. You can also filter the list by image name or tag, sort by creation time, and format the output to display only the information that is relevant to you.
	
This command is a fundamental tool for Docker users, providing the insights needed to efficiently manage Docker images on your system.`

	imageColumns = "REPOSITORY\tTAG\tIMAGE ID\tCREATED\tSIZE"

	imageOutputFormat = "%s\t%s\t%s\t%s\t%s\n"

	imageOutputQuietFormat = "%s\n"

	imageIDLength = 12
)
