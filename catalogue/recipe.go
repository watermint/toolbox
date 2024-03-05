package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	infra_recipe_rc_recipe "github.com/watermint/toolbox/infra/recipe/rc_recipe"
	recipe "github.com/watermint/toolbox/recipe"
	recipeconfigauth "github.com/watermint/toolbox/recipe/config/auth"
	recipeconfigfeature "github.com/watermint/toolbox/recipe/config/feature"
	recipedev "github.com/watermint/toolbox/recipe/dev"
	recipedevbenchmark "github.com/watermint/toolbox/recipe/dev/benchmark"
	recipedevbuild "github.com/watermint/toolbox/recipe/dev/build"
	recipedevciartifact "github.com/watermint/toolbox/recipe/dev/ci/artifact"
	recipedevciauth "github.com/watermint/toolbox/recipe/dev/ci/auth"
	recipedevdiag "github.com/watermint/toolbox/recipe/dev/diag"
	recipedevkvs "github.com/watermint/toolbox/recipe/dev/kvs"
	recipedevlifecycle "github.com/watermint/toolbox/recipe/dev/lifecycle"
	recipedevmodule "github.com/watermint/toolbox/recipe/dev/module"
	recipedevplaceholder "github.com/watermint/toolbox/recipe/dev/placeholder"
	recipedevrelease "github.com/watermint/toolbox/recipe/dev/release"
	recipedevreplay "github.com/watermint/toolbox/recipe/dev/replay"
	recipedevspec "github.com/watermint/toolbox/recipe/dev/spec"
	recipedevtest "github.com/watermint/toolbox/recipe/dev/test"
	recipedevtestauth "github.com/watermint/toolbox/recipe/dev/test/auth"
	recipedevtestsetup "github.com/watermint/toolbox/recipe/dev/test/setup"
	recipedevutil "github.com/watermint/toolbox/recipe/dev/util"
	recipedevutilimage "github.com/watermint/toolbox/recipe/dev/util/image"
	recipelogcat "github.com/watermint/toolbox/recipe/log/cat"
	recipelogjob "github.com/watermint/toolbox/recipe/log/job"
	recipeteamspaceasadminfile "github.com/watermint/toolbox/recipe/teamspace/asadmin/file"
	recipeteamspaceasadminfolder "github.com/watermint/toolbox/recipe/teamspace/asadmin/folder"
	recipeteamspaceasadminmember "github.com/watermint/toolbox/recipe/teamspace/asadmin/member"
	recipeteamspacefile "github.com/watermint/toolbox/recipe/teamspace/file"
	recipeutilarchive "github.com/watermint/toolbox/recipe/util/archive"
	recipeutilcert "github.com/watermint/toolbox/recipe/util/cert"
	recipeutildatabase "github.com/watermint/toolbox/recipe/util/database"
	recipeutildate "github.com/watermint/toolbox/recipe/util/date"
	recipeutildatetime "github.com/watermint/toolbox/recipe/util/datetime"
	recipeutildecode "github.com/watermint/toolbox/recipe/util/decode"
	recipeutildesktop "github.com/watermint/toolbox/recipe/util/desktop"
	recipeutildesktopdisplay "github.com/watermint/toolbox/recipe/util/desktop/display"
	recipeutildesktopscreenshot "github.com/watermint/toolbox/recipe/util/desktop/screenshot"
	recipeutilencode "github.com/watermint/toolbox/recipe/util/encode"
	recipeutilfile "github.com/watermint/toolbox/recipe/util/file"
	recipeutilgit "github.com/watermint/toolbox/recipe/util/git"
	recipeutilimage "github.com/watermint/toolbox/recipe/util/image"
	recipeutilmonitor "github.com/watermint/toolbox/recipe/util/monitor"
	recipeutilnet "github.com/watermint/toolbox/recipe/util/net"
	recipeutilqrcode "github.com/watermint/toolbox/recipe/util/qrcode"
	recipeutilrelease "github.com/watermint/toolbox/recipe/util/release"
	recipeutiltableformat "github.com/watermint/toolbox/recipe/util/table/format"
	recipeutiltextcase "github.com/watermint/toolbox/recipe/util/text/case"
	recipeutiltextencoding "github.com/watermint/toolbox/recipe/util/text/encoding"
	recipeutiltextnlpenglish "github.com/watermint/toolbox/recipe/util/text/nlp/english"
	recipeutiltextnlpjapanese "github.com/watermint/toolbox/recipe/util/text/nlp/japanese"
	recipeutiltidymove "github.com/watermint/toolbox/recipe/util/tidy/move"
	recipeutiltidypack "github.com/watermint/toolbox/recipe/util/tidy/pack"
	recipeutiltime "github.com/watermint/toolbox/recipe/util/time"
	recipeutilunixtime "github.com/watermint/toolbox/recipe/util/unixtime"
	recipeutiluuid "github.com/watermint/toolbox/recipe/util/uuid"
	recipeutilvideosubtitles "github.com/watermint/toolbox/recipe/util/video/subtitles"
	recipeutilxlsx "github.com/watermint/toolbox/recipe/util/xlsx"
	recipeutilxlsxsheet "github.com/watermint/toolbox/recipe/util/xlsx/sheet"
)

func AutoDetectedRecipesClassic() []infra_recipe_rc_recipe.Recipe {
	return []infra_recipe_rc_recipe.Recipe{
		&recipe.License{},
		&recipe.Version{},
		&recipeconfigauth.Delete{},
		&recipeconfigauth.List{},
		&recipeconfigfeature.Disable{},
		&recipeconfigfeature.Enable{},
		&recipeconfigfeature.List{},
		&recipedev.Info{},
		&recipedevbenchmark.Local{},
		&recipedevbenchmark.Upload{},
		&recipedevbenchmark.Uploadlink{},
		&recipedevbuild.Catalogue{},
		&recipedevbuild.Doc{},
		&recipedevbuild.Info{},
		&recipedevbuild.License{},
		&recipedevbuild.Package{},
		&recipedevbuild.Preflight{},
		&recipedevbuild.Readme{},
		&recipedevciartifact.Up{},
		&recipedevciauth.Export{},
		&recipedevdiag.Endpoint{},
		&recipedevdiag.Throughput{},
		&recipedevkvs.Concurrency{},
		&recipedevkvs.Dump{},
		&recipedevlifecycle.Planchangepath{},
		&recipedevlifecycle.Planprune{},
		&recipedevmodule.List{},
		&recipedevplaceholder.Pathchange{},
		&recipedevplaceholder.Prune{},
		&recipedevrelease.Asset{},
		&recipedevrelease.Asseturl{},
		&recipedevrelease.Candidate{},
		&recipedevrelease.Doc{},
		&recipedevrelease.Publish{},
		&recipedevreplay.Approve{},
		&recipedevreplay.Bundle{},
		&recipedevreplay.Recipe{},
		&recipedevreplay.Remote{},
		&recipedevspec.Diff{},
		&recipedevspec.Doc{},
		&recipedevtest.Echo{},
		&recipedevtest.Panic{},
		&recipedevtest.Recipe{},
		&recipedevtest.Resources{},
		&recipedevtestauth.All{},
		&recipedevtestsetup.Massfiles{},
		&recipedevtestsetup.Teamsharedlink{},
		&recipedevutil.Anonymise{},
		&recipedevutil.Wait{},
		&recipedevutilimage.Jpeg{},
		&recipelogcat.Curl{},
		&recipelogcat.Job{},
		&recipelogcat.Kind{},
		&recipelogcat.Last{},
		&recipelogjob.Archive{},
		&recipelogjob.Delete{},
		&recipelogjob.List{},
		&recipelogjob.Ship{},
		&recipeteamspaceasadminfile.List{},
		&recipeteamspaceasadminfolder.Add{},
		&recipeteamspaceasadminfolder.Delete{},
		&recipeteamspaceasadminfolder.Permdelete{},
		&recipeteamspaceasadminmember.List{},
		&recipeteamspacefile.List{},
		&recipeutilarchive.Unzip{},
		&recipeutilarchive.Zip{},
		&recipeutilcert.Selfsigned{},
		&recipeutildatabase.Exec{},
		&recipeutildatabase.Query{},
		&recipeutildate.Today{},
		&recipeutildatetime.Now{},
		&recipeutildecode.Base32{},
		&recipeutildecode.Base64{},
		&recipeutildesktop.Open{},
		&recipeutildesktopdisplay.List{},
		&recipeutildesktopscreenshot.Interval{},
		&recipeutildesktopscreenshot.Snap{},
		&recipeutilencode.Base32{},
		&recipeutilencode.Base64{},
		&recipeutilfile.Hash{},
		&recipeutilgit.Clone{},
		&recipeutilimage.Exif{},
		&recipeutilimage.Placeholder{},
		&recipeutilmonitor.Client{},
		&recipeutilnet.Download{},
		&recipeutilqrcode.Create{},
		&recipeutilqrcode.Wifi{},
		&recipeutilrelease.Install{},
		&recipeutiltableformat.Xlsx{},
		&recipeutiltextcase.Down{},
		&recipeutiltextcase.Up{},
		&recipeutiltextencoding.From{},
		&recipeutiltextencoding.To{},
		&recipeutiltextnlpenglish.Entity{},
		&recipeutiltextnlpenglish.Sentence{},
		&recipeutiltextnlpenglish.Token{},
		&recipeutiltextnlpjapanese.Token{},
		&recipeutiltextnlpjapanese.Wakati{},
		&recipeutiltidymove.Dispatch{},
		&recipeutiltidymove.Simple{},
		&recipeutiltidypack.Remote{},
		&recipeutiltime.Now{},
		&recipeutilunixtime.Format{},
		&recipeutilunixtime.Now{},
		&recipeutiluuid.V4{},
		&recipeutilvideosubtitles.Optimize{},
		&recipeutilxlsx.Create{},
		&recipeutilxlsxsheet.Export{},
		&recipeutilxlsxsheet.Import{},
		&recipeutilxlsxsheet.List{},
	}
}
