import json
import os

def parse_notables():
    source_path = "data/trade_stats.json"  
    output_path = "data/notables.json"
    
    if not os.path.exists(source_path):
        print(f"Error: {source_path} not found. Please run the new curl command first.")
        return

    with open(source_path, "r", encoding="utf-8") as f:
        data = json.load(f)

    try:
        if "result" not in data:
            print("Error: The downloaded JSON is corrupt or missing the 'result' key.")
            return

        cluster_entries = []
        
        for section in data["result"]:
            if section.get("id") == "cluster":
                cluster_entries = section.get("entries", [])
                break

        if not cluster_entries:
            print("Error: Could not find the 'cluster' stat group inside trade_stats.json.")
            return

        notables_pool = []
        unique_id = 0
        
        for entry in cluster_entries:
            trade_id = entry.get("id")   
            name = entry.get("text")      
            
            if not name or not trade_id:
                continue
                
            if not trade_id.startswith("passive_notable_"):
                continue

            notables_pool.append({
                "id": unique_id,
                "name": name,
                "trade_id": trade_id
            })
            unique_id += 1

        with open(output_path, "w", encoding="utf-8") as f:
            json.dump(notables_pool, f, indent=2)

        print(f"Successfully generated {output_path} with {len(notables_pool)} entries.")

    except (KeyError, TypeError, IndexError) as e:
        print(f"CRITICAL ERROR: Structure mismatch inside trade_stats.json.")
        print(f"Details: {e}")
        return

if __name__ == "__main__":
    parse_notables()