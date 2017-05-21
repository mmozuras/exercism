class Complement
  def self.of_dna(strand)
    /[^CGTA]/ =~ strand ? '' : strand.tr('CGTA', 'GCAU')
  end

  def self.of_rna(strand)
    strand.tr('GCAU', 'CGTA')
  end
end

module BookKeeping
  VERSION = 4
end
