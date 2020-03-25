for i in {1..20}; do 
  oc process -f ../deploy/perfApp.yml IDENTIFIER=${i} | oc apply -f -
done
