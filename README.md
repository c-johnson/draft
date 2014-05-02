draft
=====

A writing / publishing system written in Golang and backed by Camlistore.

Draft is a command-line utility that handles several tasks around writing, storing, and sharing snippets of text (most often in a blog format).  It roughly breaks down along these lines:
  1) Save and sync local Markdown files from a local directory to a cloud-backed Camlistore repository.
  2) Generates a manifest file containing metadata around when your drafts were created, what categories they fall under, if they are set to publish, etc.
  3) Based on the manifest file, generates the actual HTML content of the blog post and publishes it to my personal blog at cjohnson.io/blog

Eventually, this will be separated into clients for each platform as well as a general-purpose library for handling these functions generically.