package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
    infra_recipe_rc_recipe "github.com/watermint/toolbox/infra/recipe/rc_recipe"
    ingredientig_bootstrap "github.com/watermint/toolbox/ingredient/ig_bootstrap"
    ingredientig_dropboxig_file "github.com/watermint/toolbox/ingredient/ig_dropbox/ig_file"
    ingredientig_dropboxig_teamig_namespaceig_file "github.com/watermint/toolbox/ingredient/ig_dropbox/ig_team/ig_namespace/ig_file"
    ingredientig_dropboxig_teamig_sharedlink "github.com/watermint/toolbox/ingredient/ig_dropbox/ig_team/ig_sharedlink"
    ingredientig_dropboxig_teamfolder "github.com/watermint/toolbox/ingredient/ig_dropbox/ig_teamfolder"
    ingredientig_job "github.com/watermint/toolbox/ingredient/ig_job"
    ingredientig_releaseig_homebrew "github.com/watermint/toolbox/ingredient/ig_release/ig_homebrew"
)

func AutoDetectedIngredients() []infra_recipe_rc_recipe.Recipe {
    return []infra_recipe_rc_recipe.Recipe{
        &ingredientig_bootstrap.Autodelete{},
        &ingredientig_bootstrap.Bootstrap{},
        &ingredientig_dropboxig_file.Download{},
        &ingredientig_dropboxig_file.Online{},
        &ingredientig_dropboxig_file.Upload{},
        &ingredientig_dropboxig_teamig_namespaceig_file.List{},
        &ingredientig_dropboxig_teamig_namespaceig_file.Size{},
        &ingredientig_dropboxig_teamig_sharedlink.Update{},
        &ingredientig_dropboxig_teamfolder.Replication{},
        &ingredientig_job.Delete{},
        &ingredientig_releaseig_homebrew.Formula{},
    }
}
