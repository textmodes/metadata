# textmod.es metadata

## So you want to contribute?

That's awesome! Here are some tips to get you going.

The term **slug** is used for the name of the file, without the `.yml`
extension. So e.g., the artist file for `Hellbeard` can be found in
[artist/hellbeard.yml](artist/hellbeard.yml); in his case, his slug is
`hellbeard`. Easy, right?

## Artists

The known artists on textmod.es can be found in the [artist](artist) folder,
every artist has its own YAML file with information. For artists that go by
different aliases, we use the most dominant current used alias for the file
name, the rest of the aliases go in the `aliases` list.

Possible fields on an **artist**:

| Field     | Type         | Description  			          |
| --------- | ------------ | -------------------------------------------- |
| name      | text         | (Nick)name of the artist                     |
| aliases   | list of slug | Aliases of the artist                        |
| country   | text         | ISO 3166-1 alpha-2 country code (lower case) |
| biography | markdown[1]  | Some words about the artist's history etc.   |
| social    | map of text  | Social media profiles of the artist          |

Possible **social** site tags:

| Tag        | Site                          | Value                         |
| ---------- | ----------------------------- | ----------------------------- |
| artcity    | http://artcity.bitfellas.org/ | User ID                       |
| behance    | https://www.behance.net/      | Username                      |
| csdb       | https://csdb.dk/              | Scener ID                     |
| demozoo    | https://demozoo.org/          | Scener ID                     |
| deviantart | https://www.deviantart.com/   | Username                      |
| facebook   | https://www.facebook.com/     | Profile ID or page name       |
| flickr     | https://www.flickr.com/       | Profile ID                    |
| github     | https://github.com/           | Username                      |
| google+    | https://plus.google.com/      | +Username or Profile ID       |
| instagram  | https://instagram.com/        | Username                      |
| linkedin   | https://linkedin.com/         | Username or Profile ID        |
| pinterest  | https://pinterest.com/        | Username                      |
| pouet      | https://www.pouet.net/        | User ID                       |
| twitter    | https://twitter.com/          | Twitter handle                |
| vimeo      | https://vimeo.com/            | Username                      |
| youtube    | https://youtube.com/          | Username or Channel ID        |

## Crews

Crew information pages can be found in the [crew](crew) folder.

Possible fields on a **crew**:

| Field     | Type         | Description  			          |
| --------- | ------------ | -------------------------------------------- |
| name      | text         | Full name of the crew                        |
| aliases   | list of slug | Aliases of the crew                          |
| leaders   | list of slug | Artist slugs of the leaders of the crew      |
| website   | url          | Website of the crew                          |
| about     | markdown[1]  | Some words about the crew's history etc.     |
| members   | list of slug | Permanent crew members' slugs.               |

[1]: https://help.github.com/articles/basic-writing-and-formatting-syntax/
