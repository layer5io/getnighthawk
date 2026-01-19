#!/bin/bash

# Prompt for GitHub username
read -p "Enter your GitHub username: " GITHUB_USER

# Clone the forked repository
echo "Cloning forked repository..."
git clone git@github.com:$GITHUB_USER/getnighthawk.git
cd getnighthawk || { echo "Failed to enter repo directory"; exit 1; }

# Add upstream remote
echo "Adding upstream remote..."
git remote add upstream https://github.com/layer5io/getnighthawk.git

# Verify remotes
echo "Verifying remotes..."
git remote -v

# Fetch upstream
echo "Fetching upstream..."
git fetch upstream

# Show branches
echo "Listing branches..."
git branch -va

# Checkout master and merge upstream
echo "Updating local master branch..."
git checkout master
git merge upstream/master

# Prompt for branch type and name
echo "Creating a new branch for your work..."
read -p "Enter branch type (feature/bug): " BRANCH_TYPE
read -p "Enter issue number or short name: " BRANCH_NAME

NEW_BRANCH="$BRANCH_TYPE/$GITHUB_USER/$BRANCH_NAME"

git checkout master
git branch "$NEW_BRANCH"
git checkout "$NEW_BRANCH"

echo "You are now on branch: $NEW_BRANCH"
echo "Make your changes, commit, and push when ready."

# Instructions for rebasing before PR
echo "When ready to submit a PR, run:"
echo "  git fetch upstream"
echo "  git checkout master"
echo "  git merge upstream/master"
echo "  git checkout $NEW_BRANCH"
echo "  git rebase master"
echo "Optionally squash commits with: git rebase -i master"
