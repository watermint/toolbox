package sv_sharedlink

import "net/url"

// ToDownloadUrl converts url to force download url
// https://help.dropbox.com/share/force-download
func ToDownloadUrl(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("dl", "1")

	dl := u.Scheme + "://" + u.User.String() + u.Host + u.Path + "?" + q.Encode()

	return dl, nil
}
