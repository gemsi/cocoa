#!/bin/bash

# interaction
python3 -m scripts.responder

# widgets
python3 -m scripts.text_view
python3 -m scripts.scroll_view
python3 -m scripts.text
python3 -m scripts.view
python3 -m scripts.control
python3 -m scripts.text_container
python3 -m scripts.progress_indicator
python3 -m scripts.event
python3 -m scripts.responder
python3 -m scripts.window
python3 -m scripts.text_field
python3 -m scripts.secure_text_field
python3 -m scripts.button
python3 -m scripts.split_view
python3 -m scripts.tab_view_item
python3 -m scripts.tab_view

echo 'format go code...'
go fmt ./appkit