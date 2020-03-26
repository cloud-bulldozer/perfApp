for i in {1..20}; do 
  oc process -f ../deploy/perf-app.yml IDENTIFIER=${i} | oc apply -f -
done
