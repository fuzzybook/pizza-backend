package models

const version = "v0.0.0"

type Language struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Flag  string `json:"flag"`
}

const it = `<svg xmlns='http://www.w3.org/2000/svg' id='flag-icons-it' viewBox='0 0 640 480'>
<g fill-rule='evenodd' stroke-width='1pt'>
  <path fill='#fff' d='M0 0h640v480H0z'/>
  <path fill='#009246' d='M0 0h213.3v480H0z'/>
  <path fill='#ce2b37' d='M426.7 0H640v480H426.7z'/>
</g>
</svg>`

const de = `<svg xmlns="http://www.w3.org/2000/svg" id="flag-icons-de" viewBox="0 0 640 480">
<path fill="#ffce00" d="M0 320h640v160H0z"/>
<path d="M0 0h640v160H0z"/>
<path fill="#d00" d="M0 160h640v160H0z"/>
</svg>`

const en = `<svg xmlns='http://www.w3.org/2000/svg' id='flag-icons-gb' viewBox='0 0 640 480'>
<path fill='#012169' d='M0 0h640v480H0z'/>
<path fill='#FFF' d='m75 0 244 181L562 0h78v62L400 241l240 178v61h-80L320 301 81 480H0v-60l239-178L0 64V0h75z'/>
<path fill='#C8102E' d='m424 281 216 159v40L369 281h55zm-184 20 6 35L54 480H0l240-179zM640 0v3L391 191l2-44L590 0h50zM0 0l239 176h-60L0 42V0z'/>
<path fill='#FFF' d='M241 0v480h160V0H241zM0 160v160h640V160H0z'/>
<path fill='#C8102E' d='M0 193v96h640v-96H0zM273 0v480h96V0h-96z'/>
</svg>`

const fr = `<svg xmlns='http://www.w3.org/2000/svg' id='flag-icons-fr' viewBox='0 0 640 480'>
<path fill='#fff' d='M0 0h640v480H0z'/>
<path fill='#000091' d='M0 0h213.3v480H0z'/>
<path fill='#e1000f' d='M426.7 0H640v480H426.7z'/>
</svg>`

var languages = []Language{
	{Value: "en", Label: "English", Flag: en},
	{Value: "it", Label: "Italiano", Flag: it},
	{Value: "de", Label: "Deutsch", Flag: de},
	{Value: "fr", Label: "Français", Flag: fr},
}

var roles = []string{
	"admin",
	"user",
}

// controllers
type SystemInfoVersion struct {
	Version string `json:"version"`
}

type SystemInfo struct {
	Version   string     `json:"version"`
	Roles     []string   `json:"roles"`
	Avatars   string     `json:"avatars"`
	Languages []Language `json:"languages"`
}

func GetSystem() SystemInfo {
	var system = SystemInfo{
		Version:   version,
		Roles:     roles,
		Avatars:   "",
		Languages: languages,
	}
	return system
}
