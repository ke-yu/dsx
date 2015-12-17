import argparse
import os

def main():
    parser = argparse.ArgumentParser(description="Dump Identifier Start")
    parser.add_argument('-f', '--data_file', required=True, help='unicode data file')
    parser.add_argument('-p', '--part', action='store_true', default=False, help='if include identifier part') 
    args = parser.parse_args()

    unicode_data = args.data_file
    if unicode_data is None:
        print('Need to provide unicode data path')
        return

    # Uppercase letter (Lu)
    # Lowercase letter (Ll)
    # Titlecase letter (Lt)
    # Modifier letter (Lm)
    # Other letter (lo)
    # Letter number (Nl)
    #
    # Non-spacing mark (Mn)
    # Combining spacing mark (Mc)
    # Decimal number (Nd)
    # Connector punctation (Pc)
    # ZWNJ, ZWJ
    if args.part:
        categories = set(["Lu", "Ll", "Lt", "Lm", "Lo", "Nl", "Mn", "Mc", "Nd", "Pc", "ZERO WIDTH NON-JOINER", "ZERO WIDTH JOINER"])
    else:
        categories = set(["Lu", "Ll", "Lt", "Lm", "Lo", "Nl"])

    with open(unicode_data, 'r') as data_file:
        for line in data_file:
            fields = line.split(";")
            for field in fields:
                if field in categories:
                    decimal = int(fields[0], 16)
                    print(str(decimal) + ', ', end="")
                    break

if __name__ == "__main__":
    main()
