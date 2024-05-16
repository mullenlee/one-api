curl --request POST \
     --url http://localhost:18080/v1/chat/completions \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     -H "Authorization: Bearer sk-fD8uljCQTMwFCR6HAeAeCfF8FcAf479e95F7F157B87eB5Ee" \
     --data '
{
  "model": "llama-3-sonar-small-32k-chat",
  "messages": [
    {
      "role": "system",
      "content": "Be precise and concise."
    },
    {
      "role": "user",
      "content": "How many stars are there in our galaxy?"
    }
  ]
}
'