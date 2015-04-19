# Bowler - A GO build tool

## Bowlerfile

All options for building your project are set in the **Bowlerfile**. It should be located in the project root.
See example Bowlerfile:

	{
		"project": {
			"name": "PGProxy",
			"description": "A simple transparent PGP encryption proxy",
			"package": "git.1vh.de/maximilian.pachl/pgpproxy"
		},

		"go": {
			"min-version": "1.4.2"
		},

		"assets": [ "assets/" ]
	}

The *project* and *go* section is obligatory, while the *assets* section can be skipped.


## Assets
Bowler uses go-bindata to build your assets. See [https://github.com/jteeuwen/go-bindata] for further information. 

**IMPORTANT: Bowler creates a file called assets__.go in your project root, so make shure your project does not include such file.**


## Usage

To build your project you can use the *build* command:

	$: bowler build

Cleaning up a project is as easy as typing:

	$: bowler clean


## Tips

Your .gitignore file should exclude the **.bowler/** and **bin/** directory plus the **assets__.go** file from commits. Theses files are all artifacts which can easily generated by using the *build* command.
