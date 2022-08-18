---
layout: contributor
title: Recipe framework
lang: en
---

# Overview

A recipe is a definition and implementation of a command.
All recipe must implement functions of the interface `Recipe` of `github.com/watermint/toolbox/recipe/rc_recipe/recipe`.

```go
type Preset interface {
  Preset()
}

type Recipe interface {
Preset
Exec(c app_control.Control) error
Test(c app_control.Control) error
}
```

The function `Preset()` is the definition of default values, report model definition and input model definitions. The
framework will automatically call `Preset()` to prepare a Recipe instance before call `Exec()` or `Test()`.

# Initialization sequence

For the console based commands, commands will be launched by Bootstrap then Recipe and those RecipeValues will be
initialized through Repository. At RecipeValue spin up, if a RecipeValue is connection. Then ask for authentication to
the user. Or if it is Report RecipeValue, then opens report automatically.

{% raw %}
<div class="mermaid">
sequenceDiagram
  autonumber
  Bootstrap ->> rc_exec.DoSpec: Launch recipe
  activate rc_exec.DoSpec 
  rc_exec.DoSpec ->> rc_recipe.Repository: Retrieve recipe spec
  activate rc_recipe.Repository
  rc_recipe.Repository ->> Recipe: Instantiate
  activate Recipe
  rc_recipe.Repository ->> Recipe: SpinUp
  loop SpinUpValues
    rc_recipe.Repository ->> RecipeValue: SpinUp recipe value
  end
  rc_exec.DoSpec ->> Recipe: Execute
  rc_recipe.Repository ->> Recipe: SpinDown
  loop SpinUpValues
    rc_recipe.Repository ->> RecipeValue: SpinDown recipe value
  end
  deactivate Recipe
  deactivate rc_recipe.Repository
  deactivate rc_exec.DoSpec
</div>
{% endraw %}

## Planned change in web based commands

To achieve web based system. The initialization sequence need to be changed. For example, the console version
sequentially ask for authentication or reads input files one by one. But on the web system. UI need to spin up values
individually. For example, if the user need to connect to Dropbox account, then UI should launch authorization sequence
before entire SpinUp of the Recipe.

{% raw %}
<div class="mermaid">
sequenceDiagram
    autonumber
    User ->> WebUI: Initiate 
    activate WebUI
      WebUI ->> Repository: Retrieve Recipe spec
      activate Repository
        WebUI ->> RecipeValue[Conn]: Retrieve connection spec
        WebUI ->> User: Render Recipe options
      deactivate Repository
    deactivate WebUI

    loop All connections
        User ->> WebUI: Choose "Connect" to Dropbox
        activate WebUI
            WebUI ->> RecipeValue[Conn]: Start OAuth Sequence
            activate RecipeValue[Conn]
            RecipeValue[Conn] ->> WebUI: Return OAuth Sequnence context
            deactivate RecipeValue[Conn]
            WebUI ->> User: Render Auth dialogue
            User ->> WebUI: Proceed
            WebUI ->> Dropbox: Start OAuth sequence
            activate Dropbox
                Dropbox ->> User: Redirect
                User ->> Dropbox: Authorize
                Dropbox ->> WebUI: Redirect
            deactivate Dropbox
            WebUI ->> RecipeValue[Conn]: Complete OAuth sequence
            activate RecipeValue[Conn]
                RecipeValue[Conn] ->> AuthDatabase: Encrypt/Store OAuth
            deactivate RecipeValue[Conn]

        deactivate WebUI
    end

    User ->> WebUI: Execute
    activate WebUI
        WebUI ->> Repository: Execute
        activate Repository
            Repository ->> Recipe: Instantiate Recipe
            loop All RecipeValues
                Repository ->> RecipeValue[Conn]: Inject pre-configured values
                RecipeValue[Conn] ->> AuthDatabase: Retrieve data
            end
            Repository ->> Recipe: Execute
        deactivate Repository
    deactivate WebUI

</div>
{% endraw %}