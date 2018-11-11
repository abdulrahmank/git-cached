# git-cached

A simple git helper library to cache all the files and artifacts ignored in git using .gitignore.

Why should we cache them?

**Scenario 1:**

Working in a big project, which requires more than 5 mins to just build.
Switching back and forth between branches *without making any changes*, will consume lot of unnecessary time, which can be avoided if those build
files are cached.

**Scenario 2:**

A Project that requires changes in meta files when a change to the file system is made.

For example:
Branch1 contains file1, file2, file3, file4.
Branch2 contains also file1, file2, file3, file4.

If developer1 adds file5 to branch2, if developer2 is working on branch1 and needs to frequently checkout branch2, 
then on each checkout developer2 needs to regenerate the meta file to even develop. This happens frequently with iOS app
development where ".xcworkspace" needs to be updated when new files are added.

And there can be many other scenarios when we follow [Git-flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow "Gitflow").
This tool will be less useful for those following [Trunk based](https://trunkbaseddevelopment.com/ "Trunk based development").

**Usage**

When checking out to a new branch, run the following command to cache all the ignored files corresponding to a commit hash.

```
gitc cache
```

To apply the cache back, run the following command:

```
gitc apply
```

Note: The directory from which commands are run should be the root dir (containing .gitignore file).
If you want to run from a non root dir, please include the directory path as an argument to the command.

**WIP**

1. Calling the commands using hooks or some other mechanism in order to cache/apply the cache automatically on checking out to and from a branch.   