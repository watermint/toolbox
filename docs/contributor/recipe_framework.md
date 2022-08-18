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

The function `Preset()` is the definition of default values, report model definition and input model definitions.
The framework will automatically call `Preset()` to prepare a Recipe instance before call `Exec()` or `Test()`.

# Initialization sequence

{% raw %}
<div class="mermaid">
       graph TD 
        A[Client] --> B[Load Balancer] 
        B --> C[Server01] 
        B --> D[Server02]
</div>
{% endraw %}
