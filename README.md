# Hornbill Role Finder

The utility provides a quick and easy method of finding Roles on an instance that contains System and/or Application rights

## Execution

There are a number of MANDATORY command line parameters when running this tool:

- **instance** : The instance name (case sensitive)
- **u** : The username to use when authenticating against the instance (admin rights required)
- **p** : The password to use when authenticating the above user
- **right** : The right name to search for. This can be a partial or full right name

## Example Commmand Line Entry

To search all roles for the right `administerServiceDesk`:

- `rightFinder.exe -u=admin -p=password -instance=instanceid -right=administerServiceDesk`

To search for all roles for rights containing the string `create`

- `rightFinder.exe -u=admin -p=password -instance=instanceid -right=create`

## Example Results

Note the rows returned contain the Role name, followed by the app who owns the role (if applicable), then group & right:

- `Role [Service Desk Admin] in Application [com.hornbill.servicemanager] contains App Right [a.administerServiceDesk]`
- `Role [A Reserved Copy] in Application [com.hornbill.servicemanager] contains System Right [a.createTimer]`
- `Role [Admin Role] contains System Right [a.createRole]`

Rows that are returned in BLUE are for detected system rights, and those that are returned in GREEN are for detected application rights