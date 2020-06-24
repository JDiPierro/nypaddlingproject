# NY Paddling Project

The New York Paddling group is trying to update the information on paddling.com's Go Paddling page/app with as much information as we can on the various launch points in Upstate New York and the surrounding area.

I was able to use paddling.com's Algolia-powered search to build a list of over 700 launch points in NY and break them down by county.

This application is going to be our tool to track who's working on what, what's left to work on, and who the top contributors are.

# Idea List
- DONE: Sign in with Facebook
    - Limit contributions to members of the Upstate NY Kayaking group
- DONE: Display the list of locations, searchable by County
- Allow users to "Claim" a location that they plan to scout
    - Add checkboxes for "Updated Description" and "Added Photos" so reviewers know what to look for. 
- Once a user has updated an entry they mark it Complete - Pending Verification
- That puts it in a queue that admins have access to
- An admin verifies the update was real and high quality.
- The admin awards the user with points based on the quality of the report (5, 10, 15?)
- Display a leaderboard of who has the most points
- Add the location scraping code to this repo.
    - DONE: Pull in the Paddling.com location ID and slug
- Have a "Tutorial" page that uses fake data to walk users through claiming a location and submitting their pending updates.

# Data models:

## User
- ID
- FB auth info
- Role (User/Admin)
- Total Points

## Location
- ID
- Title
- County
- Created At (Paddling)
- Updated At (Paddling)
- Description Length
- Number of Photos
- Link to the page

## Location Claim
Having this as a separate document will allow multiple people to claim a location in case someone goes dark. Having status and points awarded on here will be good for tracking.
- User ID           # Index
- Location ID       # Index
- Created At
- Updated At
- Comment
- Status (Claimed, Pending Review, Verified)    # Index
- Points Awarded
- Reason (Admin field for rejection reason/encouraging comment)

# Pages

## Locations
Displays a DataTable of locations.
Could have the "Claim" button right in the datatable

## My Claims
Shows all locations that a user has claimed and their current status. Allows them to edit their claim's comment or status.

## Leaderboard
Shows users with the most points

## Admin Verification Queue
Only visible to people marked as Admin
Shows all claims that are Pending Verification
Admin can Verify or Reject a claim.
When verifying they choose how many points to award based on quality (Good, Great, Excellent)
