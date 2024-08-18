# Can I Commute? 🚝 🚎 🚋

A tiny tool to grab Google Maps commute times for a given origin to a set of predefined destinations.

## Purpose

I wrote this little tool to help while house hunting. Some house listing websites let you view properties by commute time but the majority don't let you specify multiple destinations. I wanted an easily way to check the approximate public transit times from a prospective new house to my office, my wife's office and our parents houses.

## How does it work?

canicommute uses the [Google Maps Distance Matrix API](https://developers.google.com/maps/documentation/distance-matrix/distance-matrix) to get the public transit times for a single origin (the prospective new house) to a number of destinations. It uses a configurable arrival time, and sets the date to the next week day (or current week day, if today is a week day).

It won't show you the transit options (so unfortunately might suggest a time based on a transit provider for which you don't have a travel card, or for a route you might find sub-optimal) but is a handy way to filter out really unsuitable properties before manually diving deeper on the transit options of potentially viable locations.

## Configuration

The config.sample.yaml file is fairly self explanatory.

Of note are the following fields:
- `auto_suffix`: allowing you to specify an automatic suffix to append to the origin location (useful for improving the reliability of geocoding).
- `api_key`: your Google Maps API key. The API key must have the Distance Matrix API enabled.
