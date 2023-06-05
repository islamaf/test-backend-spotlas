SELECT
  name,
  substring(website from '(?:https?://)?([^/?]+)') AS domain,
  COUNT(*) AS shared_domains
FROM
  "MY_TABLE"
WHERE website IS NOT NULL
GROUP BY
  name,
  domain
HAVING
  COUNT(*) > 1;
