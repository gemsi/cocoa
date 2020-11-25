#!/bin/bash

# interaction
python3 responder.py

# widgets
python3 text_view.py
python3 scroll_view.py
python3 text.py
python3 view.py
python3 control.py
python3 text_container.py
python3 progress_indicator.py
python3 event.py
python3 responder.py
python3 window.py
python3 text_field.py
python3 secure_text_field.py
python3 button.py
python3 split_view.py
python3 tab_view_item.py
python3 tab_view.py
python3 stack_view.py
python3 layout_guide.py

echo 'format go code...'
cd ..&& go fmt ./appkit