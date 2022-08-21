# Project Config

`gd` can load project specific configurations by loading a project configuration file. 
This file is called `project.toml`.

`gd` can load project configurations from following sources. 
It searches for a `project.toml` file the directory hierarchy up.

Here an example. Based on this folder structure (`monorepo-with-game`).

<details><summary>example project tree</summary>
<p>

```
monorepo-with-game/
|-- .github
|   |-- ISSUE_TEMPLATES
|   `-- workflows
|-- game
|   |-- .gitignore
|   |-- .import
|   |-- addons
|   |-- assets
|   |   |-- fonts
|   |   |-- icons
|   |   |-- music
|   |   |-- sfx
|   |   |-- sprites
|   |   `-- videos
|   |-- default_env.tres
|   |-- icon.png
|   |-- icon.png.import
|   |-- project.godot
|   |-- project.toml
|   `-- scenes
|-- marketing-functions
|   |-- .gitignore
|   |-- package.json
|   `-- src
`-- website
    |-- .gitignore
    |-- README.md
    |-- components
    |-- package.json
    |-- public
    |-- src
    `-- styles
```

</p>
</details>

We have a repository with a game folder and other folders for others aspects of our game. 
E.g. `marketing-functions` and a `website` for the game (maybe a github-page).

If the `current working directory` is `game/assets/fonts` and we execute a project related
`gd` command. The command searches and find the `game/project.toml` file.
It searches in following locations:

- `game/assets/fonts/`
- `game/assets/`
- `game/` (and will find the `project.toml` here)

> :warning: The `project.toml` file may get validated by the executed `gd` command. If the
> `project.toml` file is not valid the command will file with an error.

If the `current working directory` is `website/src` and we execute a project related
`gd` command. The command searches and won't find any `project.toml`.
It searches in following locations:

- `website/src`
- `website/`
- `./` (project root)
- `./..`
- `./../../` (til the root of the volume)

<div 
  class="warning" 
  style="background-color:#1b4c8e; color:#69337A; border-left: solid #5a8fd5 4px; border-radius: 4px; padding:0.7em;"
>

> If the executed `gd` command would change the `project.toml` file like installing a dependency
> and there is no `project.toml` file - the command will create a new `project.toml` file and
> add the expected changes from the command.

</div>
